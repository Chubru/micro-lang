// Code generated from MicroLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // MicroLang

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type MicroLangParser struct {
	*antlr.BaseParser
}

var MicroLangParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func microlangParserInit() {
	staticData := &MicroLangParserStaticData
	staticData.LiteralNames = []string{
		"", "';'", "'while'", "'{'", "'}'", "'if'", "'*'", "'/'", "'+'", "'-'",
		"'('", "')'", "'='", "'>'", "'<'", "'=='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "MULT", "DIV", "PLUS", "MINUS", "LPAREN", "RPAREN",
		"ASSIGN", "GREAT", "LESS", "EQUAL", "ID", "INT", "WS", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"prog", "stat", "expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 20, 74, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 1, 0, 5, 0, 8, 8, 0,
		10, 0, 12, 0, 11, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 29, 8, 1, 10, 1,
		12, 1, 32, 9, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1,
		42, 8, 1, 10, 1, 12, 1, 45, 9, 1, 1, 1, 1, 1, 3, 1, 49, 8, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 58, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 69, 8, 2, 10, 2, 12, 2, 72, 9, 2, 1,
		2, 0, 1, 4, 3, 0, 2, 4, 0, 3, 1, 0, 13, 15, 1, 0, 6, 7, 1, 0, 8, 9, 81,
		0, 9, 1, 0, 0, 0, 2, 48, 1, 0, 0, 0, 4, 57, 1, 0, 0, 0, 6, 8, 3, 2, 1,
		0, 7, 6, 1, 0, 0, 0, 8, 11, 1, 0, 0, 0, 9, 7, 1, 0, 0, 0, 9, 10, 1, 0,
		0, 0, 10, 12, 1, 0, 0, 0, 11, 9, 1, 0, 0, 0, 12, 13, 5, 0, 0, 1, 13, 1,
		1, 0, 0, 0, 14, 15, 5, 16, 0, 0, 15, 16, 5, 12, 0, 0, 16, 17, 3, 4, 2,
		0, 17, 18, 5, 1, 0, 0, 18, 49, 1, 0, 0, 0, 19, 20, 3, 4, 2, 0, 20, 21,
		5, 1, 0, 0, 21, 49, 1, 0, 0, 0, 22, 23, 5, 2, 0, 0, 23, 24, 5, 10, 0, 0,
		24, 25, 3, 4, 2, 0, 25, 26, 5, 11, 0, 0, 26, 30, 5, 3, 0, 0, 27, 29, 3,
		2, 1, 0, 28, 27, 1, 0, 0, 0, 29, 32, 1, 0, 0, 0, 30, 28, 1, 0, 0, 0, 30,
		31, 1, 0, 0, 0, 31, 33, 1, 0, 0, 0, 32, 30, 1, 0, 0, 0, 33, 34, 5, 4, 0,
		0, 34, 49, 1, 0, 0, 0, 35, 36, 5, 5, 0, 0, 36, 37, 5, 10, 0, 0, 37, 38,
		3, 4, 2, 0, 38, 39, 5, 11, 0, 0, 39, 43, 5, 3, 0, 0, 40, 42, 3, 2, 1, 0,
		41, 40, 1, 0, 0, 0, 42, 45, 1, 0, 0, 0, 43, 41, 1, 0, 0, 0, 43, 44, 1,
		0, 0, 0, 44, 46, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0, 46, 47, 5, 4, 0, 0, 47,
		49, 1, 0, 0, 0, 48, 14, 1, 0, 0, 0, 48, 19, 1, 0, 0, 0, 48, 22, 1, 0, 0,
		0, 48, 35, 1, 0, 0, 0, 49, 3, 1, 0, 0, 0, 50, 51, 6, 2, -1, 0, 51, 58,
		5, 17, 0, 0, 52, 58, 5, 16, 0, 0, 53, 54, 5, 10, 0, 0, 54, 55, 3, 4, 2,
		0, 55, 56, 5, 11, 0, 0, 56, 58, 1, 0, 0, 0, 57, 50, 1, 0, 0, 0, 57, 52,
		1, 0, 0, 0, 57, 53, 1, 0, 0, 0, 58, 70, 1, 0, 0, 0, 59, 60, 10, 6, 0, 0,
		60, 61, 7, 0, 0, 0, 61, 69, 3, 4, 2, 7, 62, 63, 10, 5, 0, 0, 63, 64, 7,
		1, 0, 0, 64, 69, 3, 4, 2, 6, 65, 66, 10, 4, 0, 0, 66, 67, 7, 2, 0, 0, 67,
		69, 3, 4, 2, 5, 68, 59, 1, 0, 0, 0, 68, 62, 1, 0, 0, 0, 68, 65, 1, 0, 0,
		0, 69, 72, 1, 0, 0, 0, 70, 68, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 5, 1,
		0, 0, 0, 72, 70, 1, 0, 0, 0, 7, 9, 30, 43, 48, 57, 68, 70,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// MicroLangParserInit initializes any static state used to implement MicroLangParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMicroLangParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MicroLangParserInit() {
	staticData := &MicroLangParserStaticData
	staticData.once.Do(microlangParserInit)
}

// NewMicroLangParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMicroLangParser(input antlr.TokenStream) *MicroLangParser {
	MicroLangParserInit()
	this := new(MicroLangParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &MicroLangParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "MicroLang.g4"

	return this
}

// MicroLangParser tokens.
const (
	MicroLangParserEOF          = antlr.TokenEOF
	MicroLangParserT__0         = 1
	MicroLangParserT__1         = 2
	MicroLangParserT__2         = 3
	MicroLangParserT__3         = 4
	MicroLangParserT__4         = 5
	MicroLangParserMULT         = 6
	MicroLangParserDIV          = 7
	MicroLangParserPLUS         = 8
	MicroLangParserMINUS        = 9
	MicroLangParserLPAREN       = 10
	MicroLangParserRPAREN       = 11
	MicroLangParserASSIGN       = 12
	MicroLangParserGREAT        = 13
	MicroLangParserLESS         = 14
	MicroLangParserEQUAL        = 15
	MicroLangParserID           = 16
	MicroLangParserINT          = 17
	MicroLangParserWS           = 18
	MicroLangParserCOMMENT      = 19
	MicroLangParserLINE_COMMENT = 20
)

// MicroLangParser rules.
const (
	MicroLangParserRULE_prog = 0
	MicroLangParserRULE_stat = 1
	MicroLangParserRULE_expr = 2
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStat() []IStatContext
	Stat(i int) IStatContext

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MicroLangParserRULE_prog
	return p
}

func InitEmptyProgContext(p *ProgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MicroLangParserRULE_prog
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MicroLangParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) EOF() antlr.TerminalNode {
	return s.GetToken(MicroLangParserEOF, 0)
}

func (s *ProgContext) AllStat() []IStatContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatContext); ok {
			len++
		}
	}

	tst := make([]IStatContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatContext); ok {
			tst[i] = t.(IStatContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) Stat(i int) IStatContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MicroLangListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MicroLangListener); ok {
		listenerT.ExitProg(s)
	}
}

func (s *ProgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MicroLangVisitor:
		return t.VisitProg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MicroLangParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MicroLangParserRULE_prog)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(9)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&197668) != 0 {
		{
			p.SetState(6)
			p.Stat()
		}

		p.SetState(11)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(12)
		p.Match(MicroLangParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatContext is an interface to support dynamic dispatch.
type IStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStatContext differentiates from other interfaces.
	IsStatContext()
}

type StatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatContext() *StatContext {
	var p = new(StatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MicroLangParserRULE_stat
	return p
}

func InitEmptyStatContext(p *StatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MicroLangParserRULE_stat
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MicroLangParserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

func (s *StatContext) CopyAll(ctx *StatContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IfStatContext struct {
	StatContext
}

func NewIfStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfStatContext {
	var p = new(IfStatContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *IfStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MicroLangParserLPAREN, 0)
}

func (s *IfStatContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfStatContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MicroLangParserRPAREN, 0)
}

func (s *IfStatContext) AllStat() []IStatContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatContext); ok {
			len++
		}
	}

	tst := make([]IStatContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatContext); ok {
			tst[i] = t.(IStatContext)
			i++
		}
	}

	return tst
}

func (s *IfStatContext) Stat(i int) IStatContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *IfStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MicroLangListener); ok {
		listenerT.EnterIfStat(s)
	}
}

func (s *IfStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MicroLangListener); ok {
		listenerT.ExitIfStat(s)
	}
}

func (s *IfStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MicroLangVisitor:
		return t.VisitIfStat(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignStatContext struct {
	StatContext
}

func NewAssignStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignStatContext {
	var p = new(AssignStatContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *AssignStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignStatContext) ID() antlr.TerminalNode {
	return s.GetToken(MicroLangParserID, 0)
}

func (s *AssignStatContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MicroLangParserASSIGN, 0)
}

func (s *AssignStatContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssignStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MicroLangListener); ok {
		listenerT.EnterAssignStat(s)
	}
}

func (s *AssignStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MicroLangListener); ok {
		listenerT.ExitAssignStat(s)
	}
}

func (s *AssignStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MicroLangVisitor:
		return t.VisitAssignStat(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrintExprContext struct {
	StatContext
}

func NewPrintExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintExprContext {
	var p = new(PrintExprContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *PrintExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrintExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MicroLangListener); ok {
		listenerT.EnterPrintExpr(s)
	}
}

func (s *PrintExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MicroLangListener); ok {
		listenerT.ExitPrintExpr(s)
	}
}

func (s *PrintExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MicroLangVisitor:
		return t.VisitPrintExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type WhileStatContext struct {
	StatContext
}

func NewWhileStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WhileStatContext {
	var p = new(WhileStatContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *WhileStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MicroLangParserLPAREN, 0)
}

func (s *WhileStatContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *WhileStatContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MicroLangParserRPAREN, 0)
}

func (s *WhileStatContext) AllStat() []IStatContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatContext); ok {
			len++
		}
	}

	tst := make([]IStatContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatContext); ok {
			tst[i] = t.(IStatContext)
			i++
		}
	}

	return tst
}

func (s *WhileStatContext) Stat(i int) IStatContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *WhileStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MicroLangListener); ok {
		listenerT.EnterWhileStat(s)
	}
}

func (s *WhileStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MicroLangListener); ok {
		listenerT.ExitWhileStat(s)
	}
}

func (s *WhileStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MicroLangVisitor:
		return t.VisitWhileStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MicroLangParser) Stat() (localctx IStatContext) {
	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MicroLangParserRULE_stat)
	var _la int

	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAssignStatContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(14)
			p.Match(MicroLangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(15)
			p.Match(MicroLangParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(16)
			p.expr(0)
		}
		{
			p.SetState(17)
			p.Match(MicroLangParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewPrintExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(19)
			p.expr(0)
		}
		{
			p.SetState(20)
			p.Match(MicroLangParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewWhileStatContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(22)
			p.Match(MicroLangParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(23)
			p.Match(MicroLangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(24)
			p.expr(0)
		}
		{
			p.SetState(25)
			p.Match(MicroLangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(26)
			p.Match(MicroLangParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&197668) != 0 {
			{
				p.SetState(27)
				p.Stat()
			}

			p.SetState(32)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(33)
			p.Match(MicroLangParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewIfStatContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(35)
			p.Match(MicroLangParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(36)
			p.Match(MicroLangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(37)
			p.expr(0)
		}
		{
			p.SetState(38)
			p.Match(MicroLangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(39)
			p.Match(MicroLangParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&197668) != 0 {
			{
				p.SetState(40)
				p.Stat()
			}

			p.SetState(45)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(46)
			p.Match(MicroLangParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	INT() antlr.TerminalNode
	ID() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	RPAREN() antlr.TerminalNode
	GREAT() antlr.TerminalNode
	LESS() antlr.TerminalNode
	EQUAL() antlr.TerminalNode
	MULT() antlr.TerminalNode
	DIV() antlr.TerminalNode
	PLUS() antlr.TerminalNode
	MINUS() antlr.TerminalNode

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MicroLangParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MicroLangParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MicroLangParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) GetOp() antlr.Token { return s.op }

func (s *ExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprContext) INT() antlr.TerminalNode {
	return s.GetToken(MicroLangParserINT, 0)
}

func (s *ExprContext) ID() antlr.TerminalNode {
	return s.GetToken(MicroLangParserID, 0)
}

func (s *ExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MicroLangParserLPAREN, 0)
}

func (s *ExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MicroLangParserRPAREN, 0)
}

func (s *ExprContext) GREAT() antlr.TerminalNode {
	return s.GetToken(MicroLangParserGREAT, 0)
}

func (s *ExprContext) LESS() antlr.TerminalNode {
	return s.GetToken(MicroLangParserLESS, 0)
}

func (s *ExprContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(MicroLangParserEQUAL, 0)
}

func (s *ExprContext) MULT() antlr.TerminalNode {
	return s.GetToken(MicroLangParserMULT, 0)
}

func (s *ExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(MicroLangParserDIV, 0)
}

func (s *ExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(MicroLangParserPLUS, 0)
}

func (s *ExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(MicroLangParserMINUS, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MicroLangListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MicroLangListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MicroLangVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MicroLangParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *MicroLangParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, MicroLangParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MicroLangParserINT:
		{
			p.SetState(51)
			p.Match(MicroLangParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MicroLangParserID:
		{
			p.SetState(52)
			p.Match(MicroLangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MicroLangParserLPAREN:
		{
			p.SetState(53)
			p.Match(MicroLangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(54)
			p.expr(0)
		}
		{
			p.SetState(55)
			p.Match(MicroLangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(68)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MicroLangParserRULE_expr)
				p.SetState(59)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(60)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&57344) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(61)
					p.expr(7)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MicroLangParserRULE_expr)
				p.SetState(62)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(63)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == MicroLangParserMULT || _la == MicroLangParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(64)
					p.expr(6)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MicroLangParserRULE_expr)
				p.SetState(65)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(66)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == MicroLangParserPLUS || _la == MicroLangParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(67)
					p.expr(5)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *MicroLangParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *MicroLangParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
