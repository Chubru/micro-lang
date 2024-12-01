// Code generated from MicroLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // MicroLang

import "github.com/antlr4-go/antlr/v4"

type BaseMicroLangVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseMicroLangVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMicroLangVisitor) VisitAssignStat(ctx *AssignStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMicroLangVisitor) VisitPrintExpr(ctx *PrintExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMicroLangVisitor) VisitWhileStat(ctx *WhileStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMicroLangVisitor) VisitIfStat(ctx *IfStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMicroLangVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}
