package models

type Stack struct {
    items []int
    size int
}

func (s *Stack) push(value int) {
    s.items = append(s.items, value)
    s.size = s.size + 1
}

func (s *Stack) pop() (item int, ok bool) {
    defer func() {
        if v := recover(); v != nil {
            ok = false
        }
    }()
    
    popedItem := s.items[s.size - 1]
    
    s.items = s.items[:s.size - 1]
    
    return popedItem, true
}