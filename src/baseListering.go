package src

import (
	"fmt"
	"github.com/abusizhishen/ruleEngine/parser"
	"strconv"
)

type RuleEngineListening struct {
	*parser.BaseRuleListener
	input map[string]interface{}
	output map[string]interface{}
	stack []interface{}
	errors []error
}

func (r *RuleEngineListening)pop() interface{} {
	if len(r.stack) == 0 {
		panic("stack is empty")
	}

	l := len(r.stack)
	v := r.stack[l-1]
	r.stack = r.stack[:len(r.stack)-1]
	return v
}

func (r *RuleEngineListening)push(i interface{})  {
	r.stack = append(r.stack, i)
}

func (r *RuleEngineListening)setError(e error)  {
	//r.errors = append(r.errors, e)
	panic(e.Error())
}

func New(input map[string]interface{}) *RuleEngineListening {
	var r = RuleEngineListening{
		BaseRuleListener: new(parser.BaseRuleListener),
		input:            make(map[string]interface{}),
		output:           make(map[string]interface{}),
	}

	for k,v := range input{
		r.input[k] = v
	}

	return &r
}

func (r *RuleEngineListening)EnterInit(ctx *parser.InitContext)  {
	fmt.Println("EnterInit: ",ctx.GetText())
}

func (r *RuleEngineListening)ExitInit(ctx *parser.InitContext)  {
	if len(r.errors) != 0 {
		for _,e := range r.errors{
			fmt.Println(e)
		}
		return
	}

	fmt.Println("result:", r.pop())
}

func (r *RuleEngineListening)ExitBoolStatement(ctx *parser.BoolStatementContext)  {
	fmt.Println(ctx.GetText())
	right,left := r.pop(),r.pop()

	switch ctx.GetOp().GetText() {
	case "and":
		r.push(left.(bool) && right.(bool))
	case "or":
		r.push(left.(bool) || right.(bool))
	case "not":
		r.push(left.(bool) && (!right.(bool)))
	}
}

func (r *RuleEngineListening)ExitBoolValue(ctx *parser.BoolValueContext) {
	v := r.pop()
	switch v.(type) {
	case bool:
		r.push(r)
	default:
		err := fmt.Errorf("invalid variable %s, value:%v, expect bool value", ctx.GetText(), v)
		r.setError(err)
	}
}

func (r *RuleEngineListening)ExitCalculateValue(ctx *parser.CalculateValueContext) {
	v := r.pop()
	switch v.(type) {
	case bool:
		r.push(r)
	default:
		err := fmt.Errorf("invalid variable %s, value:%v, expect bool value", ctx.GetText(), v)
		r.setError(err)
	}
}

func (r *RuleEngineListening)ExitCOMPARE(ctx *parser.COMPAREContext)  {
	fmt.Println("ExitCompare: ",ctx.GetText())
	right,left := r.pop().(int),r.pop().(int)

	switch ctx.GetOp().GetText() {
	case ">=":
		r.push(left >= right)
	case ">":
		r.push(left > right)
	case "==":
		r.push(left == right)
	case "<=":
		r.push(left <= right)
	case "<":
		r.push(left < right)
	}

}

func (r *RuleEngineListening)ExitCalculateStatement(ctx *parser.CalculateStatementContext)  {
	fmt.Println("ExitCalculateStatement: ",ctx.GetText())
	right,left := r.pop().(int),r.pop().(int)

	switch ctx.GetOp().GetText() {
	case "+":
		r.push(left + right)
	case "-":
		r.push(left - right)
	case "*":
		r.push(left * right)
	case "/":
		r.push(left / right)
	}
}

func (r *RuleEngineListening)ExitIDEN(ctx *parser.IDENContext) {
	v,ok := r.input[ctx.GetText()]
	if !ok{
		err:=fmt.Errorf("invalid variable %s", ctx.GetText())
		r.setError(err)
		return
	}

	r.push(v)
}

func (r *RuleEngineListening)ExitNUM(ctx *parser.NUMContext) {
	str := ctx.GetText()
	va,err := strconv.Atoi(str)
	if err == nil{
		r.push(va)
		return
	}

	err = fmt.Errorf("invalid num %s, err:%v", str, err)
	r.setError(err)
}

func (r *RuleEngineListening)ExitBOOL(ctx *parser.BOOLContext) {
	switch ctx.GetText() {
	case "true":
		r.push(true)
	case "false":
		r.push(false)
	}
}

func (r *RuleEngineListening)ExitIDENBOOL(ctx *parser.IDENBOOLContext) {
	v,ok := r.input[ctx.GetText()]
	if !ok{
		err:=fmt.Errorf("invalid variable %s", ctx.GetText())
		r.setError(err)
		return
	}

	va,ok := v.(bool)
	if ok{
		r.push(va)
		return
	}

	err := fmt.Errorf("invalid variable %s, value:%v, expect bool value", ctx.GetText(), v)
	r.setError(err)
}