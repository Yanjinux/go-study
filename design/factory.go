/*
 * @Author: yanjun
 * @Date: 2022-01-03 16:30:30
 * @LastEditTime: 2022-01-04 00:50:25
 * @Description:
 * @FilePath: \go-study\design\factory.go
 */

package main

import "fmt"

/**
 * @description: 交通工具通用接口
 * @param {*}
 * @return {*}
 */
type transportTool interface {
	toDst(dst string) int
	desc()
}

type car struct {
}

func (c car) desc() {
	fmt.Println("i am cat")
}
func (c car) toDst(dst string) int {
	switch dst {
	case "A":
		return 4
	case "B":
		return 5
	}
	return 0
}

type airport struct {
}

func (c airport) desc() {
	fmt.Println("i am airport")
}

func (c airport) toDst(dst string) int {
	switch dst {
	case "A":
		return 1
	case "B":
		return 2
	}
	return 0
}

// 使用地方A1
func calculateTime(dst, transportMethod string) int {
	var t transportTool
	switch transportMethod {
	case "car":
		t = car{}
	case "airport":
		t = airport{}
	default:
		return 0
	}

	return t.toDst(dst)
}

// 使用地方A2 打印运输名字
func showTransportInfo(transportMethod string) {
	var t transportTool
	switch transportMethod {
	case "car":
		t = car{}
	case "airport":
		t = airport{}
	default:
		return
	}

	t.desc()
}

func factory_main() {
	calculateTime("A", "car")
	calculateTime("B", "airport")

	showTransportInfo("car")
}
