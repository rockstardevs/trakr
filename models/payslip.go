package models

type LineItemType int

const (
   INCOME LineItemType = iota
   DEDUCTION
)

// LineItem is a single line item in a pay statement.
type LineItem struct {
  Type LineItemType
  Amount float64
  Description string
}

// Heading is a group of line items.
type Heading struct {
  Name string
  Items []LineItemType
  Total float64
}
