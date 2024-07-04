package postgres

import (
	pb "community-service/generated/community"
	"fmt"
	"reflect"
	"testing"
)

func TestCreateCommunity(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		panic(err)
	}

	community := NewCommunityRepo(db)

	rescreate, err := community.CreateCommunity(&pb.CreateCommunityRequest{
		Name:        "Nur",
		Description: "Any",
		Location:    "Tashkent",
	})
	if err != nil {
		panic(err)
	}
	waitcreate := pb.CreateCommunityResponse{
		Success: true,
	}
	if !reflect.DeepEqual(rescreate, &waitcreate) {
		t.Errorf("have %v , wont %v", rescreate, &waitcreate)
	}
}

func TestGetCommunity(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		panic(err)
	}

	community := NewCommunityRepo(db)
	resget, err := community.GetCommunity(&pb.GetCommunityRequest{
		Id: "30332b3e-433a-40b2-8066-8056bfcce188",
	})
	if err != nil {
		fmt.Println(err)
	}
	waitget := pb.GetCommunityResponse{
		Id:          "30332b3e-433a-40b2-8066-8056bfcce188",
		Name:        "Nur",
		Description: "Any",
		Location:    "Tashkent",
	}
	if !reflect.DeepEqual(resget, &waitget) {
		t.Errorf("have %v , wont %v", resget, &waitget)
	}
}

func TestUpdateCommunity(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		panic(err)
	}

	community := NewCommunityRepo(db)
	resupdate, err := community.UpdateCommunity(&pb.UpdateCommunityRequest{
		Id:          "30332b3e-433a-40b2-8066-8056bfcce188",
		Name:        "NurMel",
		Description: "Any1",
		Location:    "Tashkent",
	})
	if err != nil {
		fmt.Println(err)
	}
	waitupdate := pb.UpdateCommunityResponse{
		Succses: true,
	}
	if !reflect.DeepEqual(resupdate, &waitupdate) {
		t.Errorf("have %v , wont %v", resupdate, &waitupdate)
	}
}

func TestDeleteCommunity(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		panic(err)
	}
	community := NewCommunityRepo(db)
	resdelete, err := community.DeleteCommunity(&pb.DeleteCommunityRequest{
		Id: "30332b3e-433a-40b2-8066-8056bfcce188",
	})
	if err != nil {
		fmt.Println(err)
	}
	waitdelete := pb.DeleteCommunityResponse{
		Succses: true,
	}
	if !reflect.DeepEqual(resdelete, &waitdelete) {
		t.Errorf("have %v , wont %v", resdelete, &waitdelete)
	}
}

func TestListCommunity(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		panic(err)
	}
	community := NewCommunityRepo(db)
	reslist, err := community.ListCommunities(&pb.ListCommunitiesRequest{})
	if err != nil {
		fmt.Println(err)
	}
	waitlist := pb.ListCommunitiesResponse{
		Comunitys: []*pb.Comunity{
			{Name: "NurMel",
				Description: "Any1",
				Location:    "Tashkent",
			},
		},
	}
	if !reflect.DeepEqual(reslist, &waitlist) {
		t.Errorf("have %v , wont %v", reslist, &waitlist)
	}
}

func TestJoinCommunity(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		panic(err)
	}
	community := NewCommunityRepo(db)
	resjoin, err := community.JoinCommunity(&pb.JoinCommunityRequest{
		CommunityId: "30332b3e-433a-40b2-8066-8056bfcce188",
		UserId:      "e6165863-259b-456b-af25-35eb74031ad3",
	})
	if err != nil {
		fmt.Println(err)
	}

	waitjoin := pb.JoinCommunityResponse{
		Success: true,
	}
	if !reflect.DeepEqual(resjoin, &waitjoin) {
		t.Errorf("have %v , wont %v", resjoin, &waitjoin)
	}
}

func TestLeaveCommunity(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		panic(err)
	}
	community := NewCommunityRepo(db)
	resleave, err := community.LeaveCommunity(&pb.LeaveCommunityRequest{CommunityId: "30332b3e-433a-40b2-8066-8056bfcce188"})
	if err != nil {
		fmt.Println(err)
	}
	waitleave := pb.LeaveCommunityResponse{
		Success: true,
	}
	if !reflect.DeepEqual(resleave, &waitleave) {
		t.Errorf("have %v , wont %v", resleave, &waitleave)
	}
}

func TestCreateCommunityEvent(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		panic(err)
	}
	community := NewCommunityRepo(db)
	rescreate, err := community.CreateCommunityEvent(&pb.CreateCommunityEventRequest{
		Id:          "f449259e-7515-493e-9a74-c17bba8c8612",
		ComunityId:  "30332b3e-433a-40b2-8066-8056bfcce188",
		Name:        "ANY",
		Description: "Any",
		Type:        "workshop",
		StartType:   "2022-01-01",
		EndType:     "2022-11-01",
		Location:    "Tashkent",
	})
	if err != nil {
		fmt.Println(err)
	}
	waitcreate := pb.CreateCommunityEventResponse{
		Success: true,
	}
	if !reflect.DeepEqual(rescreate, &waitcreate) {
		t.Errorf("have %v , wont %v", rescreate, &waitcreate)
	}
}

func TestListCommunityEvent(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		panic(err)
	}
	community := NewCommunityRepo(db)
	reslist, err := community.ListCommunityEvents(&pb.ListCommunityEventsRequest{
		CommunityId: "30332b3e-433a-40b2-8066-8056bfcce188",
	})
	if err != nil {
		fmt.Println(err)
	}
	waitlist := pb.ListCommunityEventsResponse{
		CommunityEvents: []*pb.CommunityEvent{
			{
				Id:          "f449259e-7515-493e-9a74-c17bba8c8612",
				Name:        "ANY",
				Description: "Any",
				Type:        "workshop",
				StartType:   "2022-01-01T00:00:00+05:00",
				EndType:     "2022-11-01T00:00:00+05:00",
				Location:    "Tashkent",
			},
		},
	}
	if !reflect.DeepEqual(reslist, &waitlist) {
		t.Errorf("have %v , wont %v", reslist, &waitlist)
	}
}

func TestCreateCommunityForumPost(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		panic(err)
	}
	community := NewCommunityRepo(db)
	rescreate, err := community.CreateCommunityForumPost(&pb.CreateCommunityForumPostRequest{
		Id:          "f449259e-7515-493e-9a74-c17bba8c8612",
		CommunityId: "30332b3e-433a-40b2-8066-8056bfcce188",
		UserId:      "e6165863-259b-456b-af25-35eb74031ad3",
		Title:       "ANY",
		Content:     "ANY",
	})
	if err != nil {
		fmt.Println(err)
	}
	waitcreate := pb.CreateCommunityForumPostRespnse{
		Success: true,
	}
	if !reflect.DeepEqual(rescreate, &waitcreate) {
		t.Errorf("have %v , wont %v", rescreate, &waitcreate)
	}
}

func TestListCommunityForumPost(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		panic(err)
	}
	community := NewCommunityRepo(db)
	reslist, err := community.ListCommunityForumPosts(&pb.ListCommunityForumPostsRequest{
		ComunityId: "30332b3e-433a-40b2-8066-8056bfcce188",
	})
	if err != nil {
		fmt.Println(err)
	}
	waitlist := pb.ListCommunityForumPostsResponse{
		ForumPosts: []*pb.ForumPost{
			{
				Id:      "f449259e-7515-493e-9a74-c17bba8c8612",
				UserId:  "e6165863-259b-456b-af25-35eb74031ad3",
				Title:   "ANY",
				Content: "ANY",
			},
		},
	}
	if !reflect.DeepEqual(reslist, &waitlist) {
		t.Errorf("have %v , wont %v", reslist, &waitlist)
	}
}

func TestAddForumPostComment(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		panic(err)
	}
	community := NewCommunityRepo(db)
	resadd, err := community.AddForumPostComment(&pb.AddForumPostCommentRequest{
		Id:      "8d657d68-5240-4961-91b0-c3eab2bcf379",
		PostId:  "f449259e-7515-493e-9a74-c17bba8c8612",
		UserId:  "e6165863-259b-456b-af25-35eb74031ad3",
		Comment: "ANY",
	})
	if err != nil {
		fmt.Println(err)
	}
	waitadd := pb.AddForumPostCommentResponse{
		Success: true,
	}
	if !reflect.DeepEqual(resadd, &waitadd) {
		t.Errorf("have %v , wont %v", resadd, &waitadd)
	}
}

func TestListForumPostComment(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		panic(err)
	}
	community := NewCommunityRepo(db)
	reslist, err := community.ListForumPostComments(&pb.ListForumPostCommentsRequest{
		PostId: "f449259e-7515-493e-9a74-c17bba8c8612",
	})
	if err != nil {
		fmt.Println(err)
	}

	waitlist := pb.ListForumPostCommentsResponse{
		ListForumPostComments: []*pb.ListForumPostComment{
			{
				Id:      "8d657d68-5240-4961-91b0-c3eab2bcf379",
				UserId:  "e6165863-259b-456b-af25-35eb74031ad3",
				Comment: "ANY",
			},
		},
	}
	if !reflect.DeepEqual(reslist, &waitlist) {
		t.Errorf("have %v , wont %v", reslist, &waitlist)
	}
}
