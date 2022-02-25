package registry
//service
type Service struct {
    Name  string   `json:"name"`
    Nodes []*Node	`json:"nodes"`
}




// Node single node service object
type Node struct {
	Id		string	`json:"id"`
	Ip		string	`json:"ip"`
	Port	int  	`json:"port"`
	Weight	int 	`json:"weight"`
}