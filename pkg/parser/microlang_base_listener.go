// Code generated from MicroLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // MicroLang

import "github.com/antlr4-go/antlr/v4"

// BaseMicroLangListener is a complete listener for a parse tree produced by MicroLangParser.
type BaseMicroLangListener struct{}

var _ MicroLangListener = &BaseMicroLangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMicroLangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMicroLangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMicroLangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMicroLangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseMicroLangListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseMicroLangListener) ExitProg(ctx *ProgContext) {}

// EnterAssignStat is called when production assignStat is entered.
func (s *BaseMicroLangListener) EnterAssignStat(ctx *AssignStatContext) {}

// ExitAssignStat is called when production assignStat is exited.
func (s *BaseMicroLangListener) ExitAssignStat(ctx *AssignStatContext) {}

// EnterPrintExpr is called when production printExpr is entered.
func (s *BaseMicroLangListener) EnterPrintExpr(ctx *PrintExprContext) {}

// ExitPrintExpr is called when production printExpr is exited.
func (s *BaseMicroLangListener) ExitPrintExpr(ctx *PrintExprContext) {}

// EnterWhileStat is called when production whileStat is entered.
func (s *BaseMicroLangListener) EnterWhileStat(ctx *WhileStatContext) {}

// ExitWhileStat is called when production whileStat is exited.
func (s *BaseMicroLangListener) ExitWhileStat(ctx *WhileStatContext) {}

// EnterIfStat is called when production ifStat is entered.
func (s *BaseMicroLangListener) EnterIfStat(ctx *IfStatContext) {}

// ExitIfStat is called when production ifStat is exited.
func (s *BaseMicroLangListener) ExitIfStat(ctx *IfStatContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseMicroLangListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseMicroLangListener) ExitExpr(ctx *ExprContext) {}
