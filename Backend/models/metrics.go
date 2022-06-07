package models

type TicketsByStatus struct {
	Open       int64 `json:"open"`
	InProgress int64 `json:"inprogress"`
	Resolved   int64 `json:"resolved"`
}

type TicketsByPriority struct {
	Low    int64 `json:"low"`
	Medium int64 `json:"medium"`
	High   int64 `json:"high"`
}

type AdminMetrics struct {
	TotalProjects     int64             `json:"total_projects"`
	TotalTickets      int64             `json:"total_tickets"`
	TicketsByStatus   TicketsByStatus   `json:"tickets_by_status"`
	TicketsByPriority TicketsByPriority `json:"tickets_by_priority"`
}

type ProjectManagerMetrics struct {
	TotalProjects     int64             `json:"total_projects"`
	TotalTickets      int64             `json:"total_tickets"`
	TicketsByStatus   TicketsByStatus   `json:"tickets_by_status"`
	TicketsByPriority TicketsByPriority `json:"tickets_by_priority"`
}

type DeveloperMetrics struct {
	TotalProjects     int64             `json:"total_projects"`
	TotalTickets      int64             `json:"total_tickets"`
	TicketsByStatus   TicketsByStatus   `json:"tickets_by_status"`
	TicketsByPriority TicketsByPriority `json:"tickets_by_priority"`
}
