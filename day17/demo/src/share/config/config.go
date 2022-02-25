package config

const (
	MysqlDSN     ="root:@(localhost:3306)/mtbsystem"   //没有密码
	Namespace    ="com.mtsystm."
	LogPath		 ="/Users/wujiahao/study&demo/GO/package/src/go_dev/day17/demo"
	Num   		 =20 //分页每次取多少
	TickingNow	 =1  //正在上映
	TickingWill  =2  //即将上映
	ActorType	 =1  //演员
	DirectorType =2  //导演
)

const (
	ServiceNameUser   ="user"
	ServiceNameFilm   ="film"
	ServiceNameComment="comment"
	ServiceCinema     ="cinem"
	ServiceNameOrder  ="order"
	ServiceNameCMS	  ="cms"
)