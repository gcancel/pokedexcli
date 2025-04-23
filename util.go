package main

import(
	"net/url"
	"strconv"
	"fmt"
)

func ParsePageLimit(uri string) (int, int, error){
	parsedURL, err := url.Parse(uri)
	if err != nil{
		return 0, 0, fmt.Errorf("Error parsing the url... %v", err)
	}
	//fmt.Printf("url data: %v\n", parsedURL.Query())

	limit, err := strconv.Atoi(parsedURL.Query()["limit"][0])
	if err != nil{
		return 0, 0, fmt.Errorf("Error parsing limit parameter... %v", err)
	}
	offset, _ := strconv.Atoi(parsedURL.Query()["offset"][0])
	if err != nil{
		return 0, 0, fmt.Errorf("Error parsing offset parameter %v", err)
	}

	return limit, offset, nil
}