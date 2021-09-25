package main

import "fmt"

type Stage interface {
	Add()
	Delete()
	Update()
	Get()
}

type context struct {
	currentState      Stage
	currentStateIndex int
	innerStates       []Stage
}

func (ctx *context) ChangeState() {
	fmt.Println("切换状态")
	ctx.currentStateIndex = ctx.currentStateIndex + 1
	ctx.currentStateIndex %= len(ctx.innerStates)
	ctx.currentState = ctx.innerStates[ctx.currentStateIndex]
}

func (ctx *context) Work() {
	ctx.currentState.Add()
	ctx.currentState.Delete()
	ctx.currentState.Update()
	ctx.currentState.Get()
}

func Init() *context {
	ctx := &context{
		innerStates: []Stage{&PreviewState{}, &EditState{}},
	}
	ctx.currentStateIndex = 0
	ctx.currentState = ctx.innerStates[ctx.currentStateIndex]
	return ctx
}

type PreviewState struct {
}

func (state *PreviewState) Add() {
	fmt.Println("[预览模式]不支持增加操作")
}

func (state *PreviewState) Delete() {
	fmt.Println("[预览模式]不支持删除操作")
}

func (state *PreviewState) Update() {
	fmt.Println("[预览模式]不支持更新操作")
}

func (state *PreviewState) Get() {
	fmt.Println("[预览模式]查阅文档")
}

type EditState struct {
}

func (state *EditState) Add() {
	fmt.Println("[编辑模式]增加文档")
}

func (state *EditState) Delete() {
	fmt.Println("[编辑模式]删除文档")
}

func (state *EditState) Update() {
	fmt.Println("[编辑模式]更新文档")
}

func (state *EditState) Get() {
	fmt.Println("[编辑模式]查阅文档")
}

func main() {
	ctx := Init()
	ctx.Work()
	ctx.ChangeState()
	ctx.Work()
}
