package postgres

import (
	pb "community-service/generated/community"
	"community-service/pkg"
	"database/sql"
	"fmt"
)

type CommunityRepo struct {
	DB *sql.DB
}

func NewCommunityRepo(db *sql.DB) *CommunityRepo  {
	return &CommunityRepo{DB: db}
}

func (c *CommunityRepo) CreateCommunity(in *pb.CreateCommunityRequest) (*pb.CreateCommunityResponse, error) {
	rows, err := c.DB.Exec(`
			INSERS INTO
			communities(
				name,
				description,
				location
				)
			VALUES(
				$1,
				$2,
				$3)
			`, in.Name, in.Description, in.Location)

	if err != nil {
		return &pb.CreateCommunityResponse{Success: false}, err
	}

	rowsAffected, err := rows.RowsAffected()

	if err != nil || rowsAffected == 0 {
		return &pb.CreateCommunityResponse{Success: false}, err
	}

	return &pb.CreateCommunityResponse{Success: true}, nil
}

func (c *CommunityRepo) GetCommunity(in *pb.GetCommunityRequest) (*pb.GetCommunityResponse, error) {

	var resp pb.GetCommunityResponse
	err := c.DB.QueryRow(`
			SELECT
				id,
				name,
				description,
				location
			FROM communities
			WHERE
				id=$1 AND deleted_at=0
			`, in.Id).Scan(&resp.Id, &resp.Name, &resp.Description, &resp.Location)

	return &resp, err
}

func (c *CommunityRepo) UpdateCommunity(in *pb.UpdateCommunityRequest) (*pb.UpdateCommunityResponse, error) {

	params := make(map[string]interface{})

	var query = "UPDATE communities SET "
	if in.Name != "" {
		query += "name = :name, "
		params["name"] = in.Name
	}
	if in.Description != "" {
		query += "Description = :Description, "
		params["Description"] = in.Description
	}
	if in.Location != "" {
		query += "Location = :Location, "
		params["Location"] = in.Location
	}

	query += "updated_at = CURRENT_TIMESTAMP WHERE id = :id AND deleted_at = 0"
	params["id"] = in.Id
	query, args := pkg.ReplaceQueryParams(query, params)

	res, err := c.DB.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return &pb.UpdateCommunityResponse{Succses: false}, fmt.Errorf("no rows affected, user with id %s not found", in.Id)
	}

	return &pb.UpdateCommunityResponse{Succses: true}, nil

}

func (c *CommunityRepo) DeleteCommunity(in *pb.DeleteCommunityRequest) (*pb.DeleteCommunityResponse, error) {
	rows, err := c.DB.Exec(`
			UPDATE
				communities
			SET de
				deleted_ad=date_part('epoch', current_timestamp)::INT 
			WHERE
				id=$1
		`, in.Id)

	if err != nil {
		return &pb.DeleteCommunityResponse{Succses: false}, err
	}

	rowsAffected, err := rows.RowsAffected()

	if err != nil || rowsAffected == 0 {
		return &pb.DeleteCommunityResponse{Succses: false}, err
	}

	return &pb.DeleteCommunityResponse{Succses: true}, nil

}

func (c *CommunityRepo) ListCommunities(in *pb.ListCommunitiesRequest) (*pb.ListCommunitiesResponse, error) {
	rows, err := c.DB.Query(`
			SELECT
				name,
				description,
				location,
			FROM
			communities
			`)

	if err != nil {
		return nil, err
	}
	var community pb.Comunity
	var communities []*pb.Comunity
	for rows.Next() {
		err = rows.Scan(&community.Name, &community.Description, &community.Location)
		if err != nil {
			return nil, err
		}

		communities = append(communities, &community)
	}

	return &pb.ListCommunitiesResponse{Comunitys: communities}, nil
}

func (c *CommunityRepo) JoinCommunity(in *pb.JoinCommunityRequest) (*pb.JoinCommunityResponse, error) {
	return nil, nil
}

func (c *CommunityRepo) LeaveCommunity(in *pb.LeaveCommunityRequest) (*pb.LeaveCommunityResponse, error) {
	return nil, nil
}

func (c *CommunityRepo) CreateCommunityEvent(in *pb.CreateCommunityEventRequest) (*pb.CreateCommunityEventResponse,error) {
	return nil, nil

}


func (c *CommunityRepo) ListCommunityEvents(in *pb.ListCommunityEventsRequest) (*pb.ListCommunityEventsResponse,error) {
	return nil, nil

}
func (c *CommunityRepo) CreateCommunityForumPost(in *pb.CreateCommunityForumPostRequest) (*pb.CreateCommunityForumPostRespnse,error) {
	return nil, nil

}

func (c *CommunityRepo) ListCommunityForumPosts(in *pb.ListCommunityForumPostsRequest) (*pb.ListCommunityForumPostsResponse,error) {
	return nil, nil

}

func (c *CommunityRepo) AddForumPostComment(in *pb.AddForumPostCommentRequest) (*pb.AddForumPostCommentResponse,error) {
	return nil, nil

}

func (c *CommunityRepo) ListForumPostComments(in *pb.ListForumPostCommentsRequest) (*pb.ListForumPostCommentsResponse,error) {
	return nil, nil

}