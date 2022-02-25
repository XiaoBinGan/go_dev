package main

import "fmt"

func main()  {
	type AdminList struct {
		Page int `json:"page"`
	}
	type Data struct {
		id int
		name string
		mobile string
		email string
	}
	type AL struct {
		d []*Data
	}
	var dar = []*Data{
		&Data{
			id:     1,
			name:   "string",
			mobile: "string",
			email:  "string",
		},&Data{
			id:     1,
			name:   "string",
			mobile: "string",
			email:  "string",
		},	&Data{
			id:     1,
			name:   "string",
			mobile: "string",
			email:  "string",
		},	&Data{
			id:     1,
			name:   "string",
			mobile: "string",
			email:  "string",
		},
	}
	var a=&AL{
		d: dar,
	}
	fmt.Printf("%#v",a.d)
}