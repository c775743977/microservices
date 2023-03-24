package controller

import (
	"context"
	"strconv"
	"fmt"

	"gin/rpc/rpcClient"
	"gin/rpc/user"
	"gin/rpc/cart"
	"gin/rpc/book"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	var data book.PageReq
	min := c.PostForm("min")
	max := c.PostForm("max")
	if min == "" && max == "" {
		min = c.Query("min")
		max = c.Query("max")
	}
	data.PageNo, _ = strconv.ParseInt(c.DefaultQuery("PageNo", "1"), 10, 64)
	data.MaxPrice, _ = strconv.ParseFloat(max, 64)
	data.MinPrice, _ = strconv.ParseFloat(min, 64)
	userClient := rpcClient.NewUserClient()
	bookClient := rpcClient.NewBookClient()
	books, err := bookClient.GetPage(context.Background(), &data)
	if err != nil {
		c.String(500, "index/index.html", "访问页面崩溃了:(  努力抢修中~")
		return
	}
	cookie, _ := c.Cookie("uuid")
	if cookie != "" {
		res, err := userClient.GetSession(context.Background(), &user.Cookie{
			Cookie : cookie,
		})
		if err != nil {
			c.HTML(400, "index/index.html", books)
			return
		}
		books.UserName = res.UserName
		c.HTML(200, "index/index.html", books)
	} else {
		c.HTML(200, "index/index.html", books)
	}
}

func MyCartHandler(c *gin.Context) {
	ctx := context.Background()
	cartClient := rpcClient.NewCartClient()
	userClient := rpcClient.NewUserClient()
	bookClient := rpcClient.NewBookClient()
	cookie, _ := c.Cookie("uuid")
	if cookie == "" {
		c.String(400, "请先登录!")
		return
	}
	sess, err := userClient.GetSession(ctx, &user.Cookie{
		Cookie : cookie,
	})
	if err != nil {
		c.String(500, "服务器内部错误")
		return
	}
	userid, _ := strconv.ParseInt(sess.UserID, 10, 64)
	cartData, err := cartClient.GetCartByUserID(ctx, &cart.CartReqByUserID{
		UserID : userid,
	})
	cartData.UserName = sess.UserName
	for _, k := range cartData.Items {
		data, _ := bookClient.GetBook(ctx, &book.BookReq{
			ID : k.BookID,
		})
		fmt.Println("book:", data)
		k.Book = &cart.BookRes{
			ID      : data.ID,
			Title   : data.Title,
			Author  : data.Author,
			Price   : data.Price,
			Sales   : data.Sales,
			Stock   : data.Stock,
			ImgPath : data.ImgPath,
		}
	}
	cartData.Amount, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", cartData.Amount), 64)
	c.HTML(200, "cart/cart.html", cartData)
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
	bookData, err := bookClient.GetBook(ctx, &book.BookReq{
		ID : bookid,
	})
	fmt.Println("bookData:", bookData)
	if err != nil {
		c.String(500, "服务器内部错误")
		return
	}
	cookie, _ := c.Cookie("uuid")
	sess := &user.Session{}
	if cookie == "" {
		username := c.Query("user")
		if username == "" {
			c.String(400, "请先登录!")
			return
		} else {
			userData, _ := userClient.GetUserByName(ctx, &user.UserName{
				Name : username,
			})
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
		return
	}
	userid, _ := strconv.ParseInt(sess.UserID, 10, 64)
	cartData, err := cartClient.GetCartByUserID(ctx, &cart.CartReqByUserID{
		UserID : userid,
	})
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

func DelItemHandler(c *gin.Context) {
	ctx := context.Background()
	bookid, _ := strconv.ParseInt(c.Query("bookID"), 10, 64)
	cartClient := rpcClient.NewCartClient()
	userClient := rpcClient.NewUserClient()
	cookie, _ := c.Cookie("uuid")
	sess, _ := userClient.GetSession(ctx, &user.Cookie{
		Cookie : cookie,
	})
	userid, _ := strconv.ParseInt(sess.UserID, 10, 64)
	cartData, _ := cartClient.GetCartByUserID(ctx, &cart.CartReqByUserID{
		UserID : userid,
	})
	_, err := cartClient.DelItem(ctx, &cart.CartItem{
		CartID : cartData.ID,
		BookID : bookid,
	})
	if err != nil {
		c.String(500, "服务器内部错误")
		return
	}
	MyCartHandler(c)
}

func AlterItemNumHandler(c *gin.Context) {
	ctx := context.Background()
	cartClient := rpcClient.NewCartClient()
	userClient := rpcClient.NewUserClient()
	bookid, _ := strconv.ParseInt(c.Query("BookID"), 10, 64)
	booknum, _ := strconv.ParseInt(c.PostForm("BookNum"), 10, 64)
	cookie, _ := c.Cookie("uuid")
	if cookie == "" {
		c.String(400, "请登录")
		return
	}
	sess, _ := userClient.GetSession(ctx, &user.Cookie{
		Cookie : cookie,
	})
	userid, _ := strconv.ParseInt(sess.UserID, 10, 64)
	cartData, _ := cartClient.GetCartByUserID(ctx, &cart.CartReqByUserID{
		UserID : userid,
	})
	item, _ := cartClient.GetCartItem(ctx, &cart.CartItem{
		CartID : cartData.ID,
		BookID : bookid,
	})
	amount := item.Amount / float64(item.Num)
	_, err := cartClient.AlterItemNum(ctx, &cart.CartItem{
		CartID : cartData.ID,
		BookID : bookid,
		Num : booknum,
		Amount : amount,
	})
	if err != nil {
		c.String(500, "服务器内部错误")
		return
	}
	MyCartHandler(c)
}

func CleanCartHandler(c *gin.Context) {
	ctx := context.Background()
	cartClient := rpcClient.NewCartClient()
	userClient := rpcClient.NewUserClient()
	cookie, _ := c.Cookie("uuid")
	if cookie == "" {
		c.String(400, "请登录")
		return
	}
	sess, _ := userClient.GetSession(ctx, &user.Cookie{
		Cookie : cookie,
	})
	userid, _ := strconv.ParseInt(sess.UserID, 10, 64)
	cartData, _ := cartClient.GetCartByUserID(ctx, &cart.CartReqByUserID{
		UserID : userid,
	})
	_, err := cartClient.CleanCartItem(ctx, &cart.CartItem{
		CartID : cartData.ID,
	})
	if err != nil {
		c.String(500, "服务器内部错误")
		return
	}
	_, err = cartClient.CleanCart(ctx, &cart.CartReqByUserID{
		UserID : userid,
	})
	MyCartHandler(c)
}