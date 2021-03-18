package api

type LevelsRequest struct {
	Owner            string
	ProductReference string
	Location         string
}

type LevelsResponse struct {
	Stock  string
	Levels map[string]int
}

type GetLocationRequest struct {
	LocationId int `json:"location_id"`
}

type GetLocationResponse struct {
	Name string `json:"name"`
}

//todo next:
//- change URL so that GET /location/<id> -> Location{}
//- do get also for owner, company, inventory
//- do get for stock
//next write consumer to set and upd stock, stock_levels and inventory
//also consume update for the underlying tables (API restrictions should protect it)
