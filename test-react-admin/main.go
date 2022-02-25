package main

import (
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "go_dev/test-react-admin/mode"
	test_jwt "go_dev/test-react-admin/test-jwt"
	"math/rand"
	"net/http"
	"time"
)

/**
the only use in login Api
*/
func randomstr(n int) string {
	rand.Seed(time.Now().UnixNano())  //初始化种子
	b := make([]byte, n)              //随机生成字符数组
	rand.Read(b)                      //整合
	rand_str := hex.EncodeToString(b) //转换为string
	fmt.Printf("%s", rand_str)        //打印
	return rand_str
}
func main() {
	router := gin.Default()
	type data struct {
		msg  string
		code int
		toke string
	}
	type User struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	type AdminList struct {
		P int `json:"page"`
	}
	type Data struct {
		Id     int64  `json:"id"`
		Name   string `json:"name"`
		Mobile string `json:"mobile"`
		Email  string `json:"email"`
	}
	type AL struct {
		Data     []*Data `json:"data"`
		PageSize int     `json:"pagesize"`
		Total    int     `json:"total"`
	}
	var dar = []*Data{
		&Data{
			Id:     1,
			Name:   "string",
			Mobile: "string",
			Email:  "string",
		}, &Data{
			Id:     2,
			Name:   "string",
			Mobile: "string",
			Email:  "string",
		}, &Data{
			Id:     3,
			Name:   "string",
			Mobile: "string",
			Email:  "string",
		}, &Data{
			Id:     4,
			Name:   "string",
			Mobile: "string",
			Email:  "string",
		}, &Data{
			Id:     5,
			Name:   "string",
			Mobile: "string",
			Email:  "string",
		}, &Data{
			Id:     6,
			Name:   "string",
			Mobile: "string",
			Email:  "string",
		}, &Data{
			Id:     7,
			Name:   "string",
			Mobile: "string",
			Email:  "string",
		}, &Data{
			Id:     8,
			Name:   "string",
			Mobile: "string",
			Email:  "string",
		}, &Data{
			Id:     9,
			Name:   "string",
			Mobile: "string",
			Email:  "string",
		}, &Data{
			Id:     10,
			Name:   "string",
			Mobile: "string",
			Email:  "string",
		}, &Data{
			Id:     11,
			Name:   "string",
			Mobile: "string",
			Email:  "string",
		}, &Data{
			Id:     12,
			Name:   "string",
			Mobile: "string",
			Email:  "string",
		}, &Data{
			Id:     13,
			Name:   "string",
			Mobile: "string",
			Email:  "string",
		},
	}
	var a = &AL{
		Data:     dar,
		PageSize: 1,
		Total:    13,
	}

	type Role struct {
		Id       int    `json:"id"`
		RoleName string `json:"role_name"`
	}
	type RData struct {
		RoleList []*Role `json:"role_list"`
		PageSize int     `json:"pagesize"`
		Total    int     `json:"total"`
	}
	var rd = []*Role{
		&Role{
			Id:       1,
			RoleName: "string",
		}, &Role{
			Id:       2,
			RoleName: "string",
		}, &Role{
			Id:       3,
			RoleName: "string",
		}, &Role{
			Id:       4,
			RoleName: "string",
		}, &Role{
			Id:       5,
			RoleName: "string",
		}, &Role{
			Id:       6,
			RoleName: "string",
		}, &Role{
			Id:       7,
			RoleName: "string",
		}, &Role{
			Id:       8,
			RoleName: "string",
		}, &Role{
			Id:       9,
			RoleName: "string",
		}, &Role{
			Id:       10,
			RoleName: "string",
		}, &Role{
			Id:       11,
			RoleName: "string",
		}, &Role{
			Id:       12,
			RoleName: "string",
		}, &Role{
			Id:       13,
			RoleName: "string",
		}, &Role{
			Id:       14,
			RoleName: "string",
		},
	}
	var rdata = &RData{
		RoleList: rd,
		PageSize: 1,
		Total:    13,
	}

	type IPermission struct {
		Id       int		`json:"id"`
		Key      int		`json:"key"`
		IsMenu   int		`json:"ismenu"`
		ParentId int	`json:"parentid"`
		Path     string		`json:"path"`
		Title    string		`json:"title"`
	}
	type IPermissionList struct {
		List []*IPermission
	}
    var list =&IPermissionList{
    	List: []*IPermission{
			&IPermission{
				Id:       1,
				Key:      1,
				IsMenu:   1,
				ParentId: 1,
				Path:     "string",
				Title:    "string",
			},
			&IPermission{
				Id:       11,
				Key:      11,
				IsMenu:   11,
				ParentId: 1,
				Path:     "string1-1",
				Title:    "string1-1",
			},
			&IPermission{
				Id:       12,
				Key:      12,
				IsMenu:   12,
				ParentId: 1,
				Path:     "string1-2",
				Title:    "string1-2",
			},
			&IPermission{
				Id:       2,
				Key:      2,
				IsMenu:   2,
				ParentId: 2,
				Path:     "string",
				Title:    "string",
			},
			&IPermission{
				Id:       12,
				Key:      12,
				IsMenu:   12,
				ParentId: 2,
				Path:     "string1-1",
				Title:    "string1-1",
			},
			&IPermission{
				Id:       13,
				Key:      13,
				IsMenu:   13,
				ParentId: 2,
				Path:     "string1-2",
				Title:    "string1-2",
			},
		},
	}
     //jwt test Decode and Encode
	//token, err := test_jwt.GenToken("123123")
    //if err!=nil{
    //	fmt.Println(err)
	//}
	//parseToken, _ := test_jwt.ParseToken(token)
	//fmt.Print("token:",token)
	//fmt.Print("token:",parseToken)
	r := router.Group("/admin")
	//jwt middleware
	r.Use(test_jwt.JWTAuthMiddleware())
	{
		r.POST("/login", func(c *gin.Context) {
			var user User
			err := c.ShouldBind(&user)
			if err!=nil{
				c.JSON(http.StatusOK,gin.H{
					"code":2001,
					"msg":"参数无效",
				})
			}
			if user.Name == "admin" && user.Password == "admin" {
				AuthToken, _ := test_jwt.GenToken(user.Name)
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "success",
					"token":  AuthToken,
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"msg":   "please input success userName or Password",
					"code":  0,
					"token": nil,
				})
			}

		})
		r.POST("/admin/list", func(c *gin.Context) {
			j := AdminList{}
			c.Bind(&j)
			fmt.Printf("json:%#v", j)
			if int(j.P) >= 1 {
				c.JSON(http.StatusOK, a)
			} else {
				c.JSON(http.StatusOK, gin.H{
					"msg":   "please input success userName or Password",
					"code":  0,
					"token": nil,
				})
			}

		})
		r.POST("/admin/rolist", func(c *gin.Context) {
			c.JSON(http.StatusOK, rdata)
		})
		r.GET("/admin/role/detail/", func(c *gin.Context) {
			//id := c.Query("roleId")
			//if "1"==id {
				c.JSON(http.StatusOK,list)
			//}
		})
		//jwt used test
		r.POST("/auth", func(c *gin.Context) {
			var user User
			err := c.ShouldBind(&user)
			if err!=nil{
				c.JSON(http.StatusOK,gin.H{
					"code":2001,
					"msg":"参数无效",
				})
			}
			if user.Name == "admin" && user.Password == "admin" {
				AuthToken, _ := test_jwt.GenToken(user.Name)
				fmt.Printf(AuthToken)
				c.JSON(http.StatusOK, gin.H{
					"code": 2000,
					"msg":  "success",
					"data": gin.H{"token": AuthToken},
				})
			}else{
				c.JSON(http.StatusOK, gin.H{
					"code": 2002,
					"msg":  "鉴权失败",
				})
			}

		})
	}

	router.Run(":8080")
}
