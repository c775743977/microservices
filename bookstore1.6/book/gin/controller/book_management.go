package controller

import (
	"context"
	"strconv"
	"fmt"

	"github.com/gin-gonic/gin"

	"gin/rpc/rpcClient"
	"gin/rpc/user"
	"gin/rpc/book"
	"gin/rpc/cart"
)

func ManageHandler(c *gin.Context) {
	ctx := context.Background()
	userClient := rpcClient.NewUserClient()
	cookie, _ := c.Cookie("uuid")
	if cookie == "" {
		c.String(400, "没有权限!")
		return
	}
	privilege, err := userClient.IsRoot(ctx, &user.Cookie{
		Cookie : cookie,
	})
	if err != nil {
		c.String(500, "服务器内部错误")
		return
	}
	if privilege.Result {
		c.HTML(200, "manager/manager.html", nil)
	} else {
		c.String(400, "无权访问")
	}
}

func BookManageHandler(c *gin.Context) {
	ctx := context.Background()
	var data book.PageReq
	data.PageNo, _ = strconv.ParseInt(c.DefaultQuery("PageNo", "1"), 10, 64)
	bookClient := rpcClient.NewBookClient()
	books, err := bookClient.GetPage(ctx, &data)
	if err != nil {
		c.String(500, "index/index.html", "访问页面崩溃了:(  努力抢修中~")
		return
	}
	userClient := rpcClient.NewUserClient()
	cookie, _ := c.Cookie("uuid")
	if cookie == "" {
		c.String(400, "没有权限!")
		return
	}
	privilege, err := userClient.IsRoot(ctx, &user.Cookie{
		Cookie : cookie,
	})
	if err != nil {
		c.String(500, "服务器内部错误")
		return
	}
	if privilege.Result {
		c.HTML(200, "manager/book_manager.html", books)
	} else {
		c.String(400, "无权访问")
	}
}

func EditBookHandler(c *gin.Context) {
	ctx := context.Background()
	bookid, _ := strconv.ParseInt(c.Query("bookID"), 10, 64)
	userClient := rpcClient.NewUserClient()
	bookClient := rpcClient.NewBookClient()
	cookie, _ := c.Cookie("uuid")
	if cookie == "" {
		c.String(400, "没有权限!")
		return
	}
	privilege, err := userClient.IsRoot(ctx, &user.Cookie{
		Cookie : cookie,
	})
	if err != nil {
		c.String(500, "服务器内部错误")
		return
	}
	if !privilege.Result {
		c.String(400, "无权访问")
		return
	}
	if bookid == 0 {
		c.HTML(200, "manager/book_edit.html", nil)
	} else {
		data, err := bookClient.GetBook(ctx, &book.BookReq{
			ID : bookid,
		})
		if err != nil {
			c.String(500, "服务器内部错误")
			return
		}
		c.HTML(200, "manager/book_edit.html", data)
	}
}

func AddOrAlterBook(c *gin.Context) { //执行添加或者修改操作
	ctx := context.Background()
	var data book.BookRes
	bookid, _ := strconv.ParseInt(c.Query("bookID"), 10, 64)
	bookClient := rpcClient.NewBookClient()
	bbook, err := bookClient.GetBook(ctx, &book.BookReq{
		ID : bookid,
	})
	if err != nil {
		c.String(500, "服务器内部错误")
		return
	}
	err = c.Bind(&data)
	data.ImgPath = "static/img/default.jpg"
	if data.Title == bbook.Title && data.Author == bbook.Author {
		_, err = bookClient.AlterBook(ctx, &data)
		BookManageHandler(c)
		return
	}
	//报错是说明提交了空白数据
	if err != nil {
		data.Err = "失败！不允许提交空白数据"
		c.HTML(400, "manager/book_edit.html", data)
		return
	}
	if data.ID == 0 {
		if CheckBook(data.Title, data.Author) { //判断图书是否存在
			data.Err = "失败！该书已存在"
			c.HTML(400, "manager/book_edit.html", data)
			return
		}
		_, err = bookClient.AddBook(ctx, &data)
		BookManageHandler(c)
	} else {
		if bbook.Title != data.Title && CheckBook(data.Title, data.Author) {
			data.Err = "失败！该书已存在"
			c.HTML(400, "manager/book_edit.html", data)
			return
		}
		_, err = bookClient.AlterBook(ctx, &data)
		BookManageHandler(c)
	}
}

func CheckBook(title string, author string) bool {
	bookClient := rpcClient.NewBookClient()
	data, err := bookClient.GetBookByName(context.Background(), &book.BookRes{
		Title : title,
	})
	fmt.Printf("data:%v\ntitle:%s\nauthor:%s\n", data, title, author)
	if err != nil {
		fmt.Println("CheckBook error:", err)
		return true
	}
	for _, k := range data.Books {
		if k.Author == author {
			return true
		}
	}
	return false
}

func DelBookHandler(c *gin.Context) {
	ctx := context.Background()
	bookid, _ := strconv.ParseInt(c.Query("bookID"), 10, 64)
	bookClient := rpcClient.NewBookClient()
	_, err := bookClient.DelBook(ctx, &book.BookReq{
		ID : bookid,
	})
	if err != nil {
		c.String(500, "服务器内部错误")
		return
	}
	BookManageHandler(c)
}

func CheckItem(ctx context.Context, bookid int64, cartid int64) bool {
	cartClient := rpcClient.NewCartClient()
	data, err := cartClient.GetCartByID(ctx, &cart.CartReqByID{
		CartID : cartid,
	})
	if err != nil {
		fmt.Println("call rpc-GetCartByID error:", err)
		return false
	}
	for _,k := range data.Items {
		if k.BookID == bookid {
			return true
		}
	}
	return false
}

func AddItemHandler(c *gin.Context) {
	ctx := context.Background()
	cartClient := rpcClient.NewCartClient()
	userClient := rpcClient.NewUserClient()
	bookClient := rpcClient.NewBookClient()
	bookid, _ := strconv.ParseInt(c.PostForm("bookId"), 10, 64)
	fmt.Println("bookid:", bookid)
	bookData, err := bookClient.GetBook(ctx, &book.BookReq{
		ID : bookid,
	})
	fmt.Println("bookData:", bookData)
	if err != nil {
		c.String(500, "服务器内部错误")
		return
	}
	cookie, _ := c.Cookie("uuid")
	fmt.Println("cookie:", cookie)
	sess := &user.Session{}
	if cookie == "" {
		fmt.Println("未能获取到登录信息")
		username := c.Query("user")
		fmt.Println("username:", username)
		if username == "" {
			c.String(400, "请先登录!")
			return
		} else {
			userData, _ := userClient.GetUserByName(ctx, &user.UserName{
				Name : username,
			})
			fmt.Println("userData:", userData)
			sess.UserID = fmt.Sprint(userData.UserID)
			sess.UserName = username
		}
	} else {
		sess, err = userClient.GetSession(ctx, &user.Cookie{
			Cookie : cookie,
		})
	}
	fmt.Println("session:", sess)
	if err != nil {
		c.String(500, "服务器内部错误")
		return
	}
	userid, _ := strconv.ParseInt(sess.UserID, 10, 64)
	cartData, err := cartClient.GetCartByUserID(ctx, &cart.CartReqByUserID{
		UserID : userid,
	})
	fmt.Println("cartData:", cartData)
	if CheckItem(ctx, bookid, cartData.ID) {
		item, _ := cartClient.GetCartItem(ctx, &cart.CartItem{
			CartID : cartData.ID,
			BookID : bookid,
		})
		_, err = cartClient.AddItemNum(ctx, &cart.CartItem{
			CartID : cartData.ID,
			BookID : bookid,
			Num : item.Num,
			Amount : item.Amount,
		})
		if err != nil {
			c.String(500, "服务器内部错误")
			return
		}
	} else {
		bookres := &cart.BookRes{
			ID      : bookData.ID,
			Title   : bookData.Title,
			Author  : bookData.Author,
			Price   : bookData.Price,
			Sales   : bookData.Sales,
			Stock   : bookData.Stock,
			ImgPath : bookData.ImgPath,
		}
		_, err = cartClient.AddItem(ctx, &cart.CartItem{
			CartID : cartData.ID,
			BookID : bookid,
			Num : 1,
			Amount : bookData.Price,
			Book : bookres,
		})
	}
	IndexHandler(c)
}