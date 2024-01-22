package main

import (
	"fmt"
	"os"
	"github.com/goplus/yap"
	"context"
	"io"
	"net/http"
	"github.com/goplus/community/internal/core"
	"github.com/goplus/community/translation"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"go.uber.org/zap"
	"golang.org/x/text/language"
	_ "github.com/joho/godotenv/autoload"
)

type community struct {
	yap.App
	community *core.Community
	trans     *translation.Engine
}
//line cmd/gopcomm/community_yap.gox:23
func (this *community) MainEntry() {
//line cmd/gopcomm/community_yap.gox:23:1
	todo := context.TODO()
//line cmd/gopcomm/community_yap.gox:24:1
	endpoint := os.Getenv("GOP_COMMUNITY_ENDPOINT")
//line cmd/gopcomm/community_yap.gox:25:1
	domain := os.Getenv("GOP_COMMUNITY_DOMAIN")
//line cmd/gopcomm/community_yap.gox:26:1
	logger, _ := zap.NewProduction()
//line cmd/gopcomm/community_yap.gox:27:1
	defer logger.Sync()
//line cmd/gopcomm/community_yap.gox:28:1
	zlog := logger.Sugar()
//line cmd/gopcomm/community_yap.gox:30:1
	this.Static__0("/")
//line cmd/gopcomm/community_yap.gox:32:1
	this.Get("/success", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:33:1
		ctx.Yap__1("2xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:36:1
	this.Get("/error", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:37:1
		ctx.Yap__1("4xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:40:1
	this.Get("/failed", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:41:1
		ctx.Yap__1("5xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:44:1
	this.Get("/p/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:45:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:46:1
		article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:47:1
		ctx.Yap__1("article", map[string]interface {
		}{"ID": id, "Title": article.Title, "Content": article.Content, "Tags": article.Tags, "Cover": article.Cover, "Ctime": article.Ctime, "User": article.User})
	})
//line cmd/gopcomm/community_yap.gox:58:1
	this.Get("/getArticle/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:59:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:60:1
		article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:61:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": article})
	})
//line cmd/gopcomm/community_yap.gox:67:1
	this.Get("/user/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:68:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:72:1
		// Get User Info
		var userClaim *casdoorsdk.Claims
//line cmd/gopcomm/community_yap.gox:73:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:74:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:75:1
			userClaim, err = this.community.GetUserClaim(token.Value)
//line cmd/gopcomm/community_yap.gox:76:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:77:1
				zlog.Error("get user claim error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:82:1
		ctx.Yap__1("user", map[string]interface {
		}{"Id": id, "CurrentUser": userClaim, "User": userClaim})
	})
//line cmd/gopcomm/community_yap.gox:89:1
	this.Get("/add", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:90:1
		ctx.Yap__1("edit", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:93:1
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:97:1
		// Get User Info
		var userClaim *casdoorsdk.Claims
//line cmd/gopcomm/community_yap.gox:98:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:99:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:100:1
			userClaim, err = this.community.GetUserClaim(token.Value)
//line cmd/gopcomm/community_yap.gox:101:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:102:1
				zlog.Error("get user claim error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:107:1
		articles, _, _ := this.community.ListArticle(todo, core.MarkBegin, 20)
//line cmd/gopcomm/community_yap.gox:108:1
		ctx.Yap__1("home", map[string]interface {
		}{"User": userClaim, "Items": articles})
	})
//line cmd/gopcomm/community_yap.gox:114:1
	this.Post("/search", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:115:1
		searchValue := ctx.Param("search")
//line cmd/gopcomm/community_yap.gox:116:1
		zlog.Infof("SearchValue: %+v", searchValue)
//line cmd/gopcomm/community_yap.gox:117:1
		if searchValue == "" {
//line cmd/gopcomm/community_yap.gox:118:1
			ctx.Json__1(map[string]interface {
			}{"code": 400, "err": "value can not be ''."})
		}
//line cmd/gopcomm/community_yap.gox:123:1
		articles, _ := this.community.SearchArticle(todo, searchValue)
//line cmd/gopcomm/community_yap.gox:124:1
		ctx.Yap__1("home", map[string][]*core.ArticleEntry{"Items": articles})
	})
//line cmd/gopcomm/community_yap.gox:129:1
	this.Get("/edit/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:130:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:131:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:132:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:138:1
		uid, err := this.community.GetUserId(token.Value)
//line cmd/gopcomm/community_yap.gox:139:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:140:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:146:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:147:1
		if id != "" {
//line cmd/gopcomm/community_yap.gox:148:1
			if
//line cmd/gopcomm/community_yap.gox:148:1
			editable, _ := this.community.CanEditable(todo, uid, id); !editable {
//line cmd/gopcomm/community_yap.gox:149:1
				zlog.Error("no permissions")
//line cmd/gopcomm/community_yap.gox:150:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:152:1
			article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:153:1
			ctx.Yap__1("edit", article)
		}
	})
//line cmd/gopcomm/community_yap.gox:157:1
	this.Get("/getTrans", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:158:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:159:1
		htmlUrl, err := this.community.TransHtmlUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:160:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:161:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:166:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:173:1
	this.Post("/commit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:175:1
		trans := ctx.Param("trans")
//line cmd/gopcomm/community_yap.gox:176:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:177:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:178:1
		htmlData := ctx.Param("html")
//line cmd/gopcomm/community_yap.gox:180:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:181:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:182:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:187:1
		uid, err := this.community.GetUserId(token.Value)
//line cmd/gopcomm/community_yap.gox:188:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:189:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:196:1
		article := &core.Article{ArticleEntry: core.ArticleEntry{ID: id, Title: ctx.Param("title"), UId: uid, Cover: ctx.Param("cover"), Tags: ctx.Param("tags")}, Content: mdData, HtmlData: htmlData}
//line cmd/gopcomm/community_yap.gox:207:1
		id, _ = this.community.PutArticle(todo, uid, trans, article)
//line cmd/gopcomm/community_yap.gox:208:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": id})
	})
//line cmd/gopcomm/community_yap.gox:216:1
	this.Post("/translate", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:218:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:219:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:220:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:225:1
		uid, err := this.community.GetUserId(token.Value)
//line cmd/gopcomm/community_yap.gox:226:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:227:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:233:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:234:1
		htmlData := ctx.Param("html")
//line cmd/gopcomm/community_yap.gox:235:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:236:1
		zlog.Info(mdData)
//line cmd/gopcomm/community_yap.gox:238:1
		transData, err := this.trans.TranslateMarkdownText(mdData, language.Chinese, language.English)
//line cmd/gopcomm/community_yap.gox:239:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:240:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:245:1
		id, _ = this.community.SaveHtml(todo, uid, htmlData, mdData, id)
//line cmd/gopcomm/community_yap.gox:246:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "id": id, "data": transData})
	})
//line cmd/gopcomm/community_yap.gox:253:1
	this.Get("/getMedia/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:254:1
		mediaId := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:256:1
		fileKey, _ := this.community.GetMediaUrl(context.Background(), mediaId)
//line cmd/gopcomm/community_yap.gox:258:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, domain+fileKey, http.StatusTemporaryRedirect)
	})
//line cmd/gopcomm/community_yap.gox:261:1
	this.Get("/getMediaUrl/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:262:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:263:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:264:1
		htmlUrl := fmt.Sprintf("%s%s", domain, fileKey)
//line cmd/gopcomm/community_yap.gox:265:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:266:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:271:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "url": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:277:1
	this.Post("/upload", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:278:1
		file, header, err := ctx.FormFile("file")
//line cmd/gopcomm/community_yap.gox:279:1
		filename := header.Filename
//line cmd/gopcomm/community_yap.gox:281:1
		ctx.ParseMultipartForm(10 << 20)
//line cmd/gopcomm/community_yap.gox:283:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:284:1
			zlog.Error("upload file error:", filename)
//line cmd/gopcomm/community_yap.gox:285:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:286:1
			return
		}
//line cmd/gopcomm/community_yap.gox:290:1
		dst, err := os.Create(filename)
//line cmd/gopcomm/community_yap.gox:291:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:292:1
			zlog.Error("create file error:", file)
//line cmd/gopcomm/community_yap.gox:293:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:294:1
			return
		}
//line cmd/gopcomm/community_yap.gox:296:1
		defer func() {
//line cmd/gopcomm/community_yap.gox:297:1
			file.Close()
//line cmd/gopcomm/community_yap.gox:298:1
			dst.Close()
//line cmd/gopcomm/community_yap.gox:299:1
			err = os.Remove(filename)
//line cmd/gopcomm/community_yap.gox:300:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:301:1
				zlog.Error("delete file error:", filename)
//line cmd/gopcomm/community_yap.gox:302:1
				return
			}
		}()
//line cmd/gopcomm/community_yap.gox:307:1
		_, err = io.Copy(dst, file)
//line cmd/gopcomm/community_yap.gox:308:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:309:1
			zlog.Error("copy file errer:", filename)
//line cmd/gopcomm/community_yap.gox:310:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:311:1
			return
		}
//line cmd/gopcomm/community_yap.gox:313:1
		bytes, err := os.ReadFile(filename)
//line cmd/gopcomm/community_yap.gox:314:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:315:1
			zlog.Error("read file errer:", filename)
//line cmd/gopcomm/community_yap.gox:316:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:317:1
			return
		}
//line cmd/gopcomm/community_yap.gox:319:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:320:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:321:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:326:1
		uid, err := this.community.GetUserId(token.Value)
//line cmd/gopcomm/community_yap.gox:327:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:328:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:335:1
		id, err := this.community.SaveMedia(context.Background(), uid, bytes)
//line cmd/gopcomm/community_yap.gox:336:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:337:1
			zlog.Error("save file", err.Error())
//line cmd/gopcomm/community_yap.gox:338:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:339:1
			return
		}
//line cmd/gopcomm/community_yap.gox:343:1
		ctx.JSON(200, id)
	})
//line cmd/gopcomm/community_yap.gox:346:1
	conf := &core.Config{}
//line cmd/gopcomm/community_yap.gox:347:1
	this.community, _ = core.New(todo, conf)
//line cmd/gopcomm/community_yap.gox:348:1
	this.trans = translation.New(os.Getenv("NIUTRANS_API_KEY"), "", "")
//line cmd/gopcomm/community_yap.gox:349:1
	core.CasdoorConfigInit()
//line cmd/gopcomm/community_yap.gox:351:1
	zlog.Info("Started in endpoint: ", endpoint)
//line cmd/gopcomm/community_yap.gox:352:1
	this.Run(endpoint)
}
func main() {
	yap.Gopt_App_Main(new(community))
}
