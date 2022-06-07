package controllers

import (
	"Backend/config"
	"Backend/models"
	"container/list"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func MetricsByRole(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	Author, CompanyCode, AuthorRole, err := authenticate(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	client := config.ClientConnection()
	ticketColl := client.Database(config.ViperEnvVariable("dbName")).Collection("tickets")
	projectColl := client.Database(config.ViperEnvVariable("dbName")).Collection("projects")

	if AuthorRole == "admin" {
		totalProjects, err := projectColl.CountDocuments(context.TODO(), bson.D{{"company_code", CompanyCode}})
		errHandler(w, err)
		totalTickets, err := ticketColl.CountDocuments(context.TODO(), bson.D{{"company_code", CompanyCode}})
		errHandler(w, err)
		totalOpenTickets, err := ticketColl.CountDocuments(context.TODO(), bson.D{{"company_code", CompanyCode}, {"status", "open"}})
		errHandler(w, err)
		totalInProgressTickets, err := ticketColl.CountDocuments(context.TODO(), bson.D{{"company_code", CompanyCode}, {"status", "in-progress"}})
		errHandler(w, err)
		totalResolvedTickets, err := ticketColl.CountDocuments(context.TODO(), bson.D{{"company_code", CompanyCode}, {"status", "resolved"}})
		errHandler(w, err)
		totalLowTickets, err := ticketColl.CountDocuments(context.TODO(), bson.D{{"company_code", CompanyCode}, {"priority", "low"}})
		errHandler(w, err)
		totalMediumTickets, err := ticketColl.CountDocuments(context.TODO(), bson.D{{"company_code", CompanyCode}, {"priority", "medium"}})
		errHandler(w, err)
		totalHighTickets, err := ticketColl.CountDocuments(context.TODO(), bson.D{{"company_code", CompanyCode}, {"priority", "high"}})
		errHandler(w, err)
		metrics := models.AdminMetrics{
			TotalProjects: totalProjects,
			TotalTickets:  totalTickets,
			TicketsByStatus: models.TicketsByStatus{
				Open:       totalOpenTickets,
				InProgress: totalInProgressTickets,
				Resolved:   totalResolvedTickets,
			},
			TicketsByPriority: models.TicketsByPriority{
				Low:    totalLowTickets,
				Medium: totalMediumTickets,
				High:   totalHighTickets,
			},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(metrics)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	} else if AuthorRole == "project-manager" {
		filter := bson.D{{"project_manager", Author}, {"company_code", CompanyCode}}
		var projects []models.Project
		var tickets []models.Ticket
		cursor, err := projectColl.Find(context.TODO(), filter)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err = cursor.All(context.TODO(), &projects); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := cursor.Close(context.TODO()); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		ticketList := list.New()
		for _, project := range projects {
			cursor, err := ticketColl.Find(context.TODO(), bson.D{{"project_id", project.Id}, {"company_code", CompanyCode}})
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			for cursor.Next(context.TODO()) {
				var ticket models.Ticket
				err = cursor.Decode(&ticket)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				ticketList.PushBack(ticket)
			}
			if err := cursor.Close(context.TODO()); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
		for e := ticketList.Front(); e != nil; e = e.Next() {
			tickets = append(tickets, e.Value.(models.Ticket))
		}

		totalProjects := len(projects)
		totalTickets := len(tickets)
		totalOpenTickets := 0
		totalInProgressTickets := 0
		totalResolvedTickets := 0
		totalLowTickets := 0
		totalMediumTickets := 0
		totalHighTickets := 0
		for _, ticket := range tickets {
			if ticket.Status == "open" {
				totalOpenTickets++
			} else if ticket.Status == "in-progress" {
				totalInProgressTickets++
			} else if ticket.Status == "resolved" {
				totalResolvedTickets++
			}

			if ticket.Priority == "low" {
				totalLowTickets++
			} else if ticket.Priority == "medium" {
				totalMediumTickets++
			} else if ticket.Priority == "high" {
				totalHighTickets++
			}
		}

		metrics := models.ProjectManagerMetrics{
			TotalProjects: int64(totalProjects),
			TotalTickets:  int64(totalTickets),
			TicketsByStatus: models.TicketsByStatus{
				Open:       int64(totalOpenTickets),
				InProgress: int64(totalInProgressTickets),
				Resolved:   int64(totalResolvedTickets),
			},
			TicketsByPriority: models.TicketsByPriority{
				Low:    int64(totalLowTickets),
				Medium: int64(totalMediumTickets),
				High:   int64(totalHighTickets),
			},
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(metrics)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return

	} else if AuthorRole == "developer" {
		totalProjectsAssigned, err := projectColl.CountDocuments(context.TODO(), bson.D{{"company_code", CompanyCode}, {"assigned_to", Author}})
		errHandler(w, err)
		totalTicketsAssigned, err := ticketColl.CountDocuments(context.TODO(), bson.D{{"company_code", CompanyCode}, {"assigned_to", Author}})
		errHandler(w, err)
		totalOpenTicketsAssigned, err := ticketColl.CountDocuments(context.TODO(), bson.D{{"company_code", CompanyCode}, {"assigned_to", Author}, {"status", "open"}})
		errHandler(w, err)
		totalInProgressTicketsAssigned, err := ticketColl.CountDocuments(context.TODO(), bson.D{{"company_code", CompanyCode}, {"assigned_to", Author}, {"status", "in-progress"}})
		errHandler(w, err)
		totalResolvedTicketsAssigned, err := ticketColl.CountDocuments(context.TODO(), bson.D{{"company_code", CompanyCode}, {"assigned_to", Author}, {"status", "resolved"}})
		errHandler(w, err)
		totalLowTicketsAssigned, err := ticketColl.CountDocuments(context.TODO(), bson.D{{"company_code", CompanyCode}, {"assigned_to", Author}, {"priority", "low"}})
		errHandler(w, err)
		totalMediumTicketsAssigned, err := ticketColl.CountDocuments(context.TODO(), bson.D{{"company_code", CompanyCode}, {"assigned_to", Author}, {"priority", "medium"}})
		errHandler(w, err)
		totalHighTicketsAssigned, err := ticketColl.CountDocuments(context.TODO(), bson.D{{"company_code", CompanyCode}, {"assigned_to", Author}, {"priority", "high"}})
		errHandler(w, err)
		metrics := models.DeveloperMetrics{
			TotalProjects: totalProjectsAssigned,
			TotalTickets:  totalTicketsAssigned,
			TicketsByStatus: models.TicketsByStatus{
				Open:       totalOpenTicketsAssigned,
				InProgress: totalInProgressTicketsAssigned,
				Resolved:   totalResolvedTicketsAssigned,
			},
			TicketsByPriority: models.TicketsByPriority{
				Low:    totalLowTicketsAssigned,
				Medium: totalMediumTicketsAssigned,
				High:   totalHighTicketsAssigned,
			},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(metrics)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}

func errHandler(w http.ResponseWriter, err error) {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
