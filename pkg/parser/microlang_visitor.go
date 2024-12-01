// Code generated from MicroLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // MicroLang

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by MicroLangParser.
type MicroLangVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MicroLangParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by MicroLangParser#assignStat.
	VisitAssignStat(ctx *AssignStatContext) interface{}

	// Visit a parse tree produced by MicroLangParser#printExpr.
	VisitPrintExpr(ctx *PrintExprContext) interface{}

	// Visit a parse tree produced by MicroLangParser#whileStat.
	VisitWhileStat(ctx *WhileStatContext) interface{}

	// Visit a parse tree produced by MicroLangParser#ifStat.
	VisitIfStat(ctx *IfStatContext) interface{}

	// Visit a parse tree produced by MicroLangParser#expr.
	VisitExpr(ctx *ExprContext) interface{}
}
