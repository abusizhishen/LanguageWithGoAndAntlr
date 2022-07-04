package src

import (
	"fmt"
	"github.com/abusizhishen/LanguageWithGoAndAntlr/parser"
	"reflect"
)

func (v *RuleEngineVisitor) VisitGetMapOrArrayValue(ctx *parser.GetMapOrArrayValueContext) interface{} {
	fmt.Println("VisitGetMapOrArrayValue:", ctx.GetText())

	ctx.Identify().Accept(v)
	for _, key := range ctx.AllMapKey() {
		fmt.Println(key)
		key.Accept(v)
		//keyType := v.pop()
		//switch keyType.(type) {
		//case string:
		//	v.pu
		//	//v.HandlerIStringValue()
		//}
		//switch key.(type) {
		//case parser.INumContext:
		//	v.handlerArray(key.(parser.INumContext))
		//case parser.IStringValueContext:
		//	v.HandlerIStringValue(key.(parser.IStringValueContext))
		//default:
		//	ty := reflect.TypeOf(key)
		//	fmt.Println(111, ty.String())
		//}

		//value := v.pop()
		//key.Accept(v)
		//valueType := v.pop()
		//switch valueType.(type) {
		//case float64:
		//	v.VisitArray()
		//case string:
		//	v.VisitMap()
		//}
	}

	return nil
}

func (v *RuleEngineVisitor) VisitMapKey(ctx *parser.MapKeyContext) interface{} {
	fmt.Println("VisitMapKey:", ctx.GetText())

	for i := 0; i < ctx.GetChildCount(); i++ {
		child := ctx.GetChildOfType(i, nil)

		ty := reflect.TypeOf(child)
		fmt.Println(111, ty.String(), ty.Name())
		child.Accept(v)

		switch child.(type) {
		case *parser.NUMContext:
			child.(*parser.NUMContext).Accept(v)
		case *parser.STRINGContext:
			child.(*parser.STRINGContext).Accept(v)
		case *parser.IDENTIFYContext:
			child.(*parser.IDENTIFYContext).Accept(v)
		default:
			ty := reflect.TypeOf(child)
			v.setError(fmt.Errorf("无效的类型:%s", ty.String()))
		}
	}

	return nil
}
