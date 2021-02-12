package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/cpartogi/wgql/graph/generated"
	"github.com/cpartogi/wgql/graph/model"
)

func (r *queryResolver) Menutype(ctx context.Context) ([]*model.MenuType, error) {
	resp, err := http.Get(os.Getenv("MENU_SVC_URL") + "menus/typelist")

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	var responseObject model.ResponseMenuType

	json.Unmarshal(body, &responseObject)

	statusCode := responseObject.StatusCode

	if statusCode != 200 {
		return nil, errors.New(responseObject.Message)
	}

	var menuTypes []*model.MenuType

	for i := 0; i < len(responseObject.Data); i++ {
		menuType := &model.MenuType{
			ID:   responseObject.Data[i].MenuTypeId,
			Name: responseObject.Data[i].MenuTypeName,
		}

		menuTypes = append(menuTypes, menuType)
	}
	return menuTypes, nil
}

func (r *queryResolver) Menudetail(ctx context.Context, menuID string) (*model.MenuDetail, error) {
	resp, err := http.Get(os.Getenv("MENU_SVC_URL") + "menu/" + menuID)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	var responseObject model.ResponseMenuDetail

	json.Unmarshal(body, &responseObject)

	statusCode := responseObject.StatusCode

	if statusCode != 200 {
		return nil, errors.New(responseObject.Message)
	}

	menuDetail := &model.MenuDetail{
		MenuId:   menuID,
		Name:     responseObject.Data.MenuName,
		Detail:   responseObject.Data.MenuDetail,
		Picture:  responseObject.Data.MenuPicture,
		Price:    responseObject.Data.MenuPrice,
		TypeName: responseObject.Data.MenuTypeName,
		WartegId: responseObject.Data.WartegId,
	}
	return menuDetail, nil
}

func (r *queryResolver) Menulist(ctx context.Context, wartegID *string, menuTypeID *string, menuName *string) ([]*model.MenuList, error) {
	resp, err := http.Get(os.Getenv("MENU_SVC_URL") + "menus/list?warteg_id=" + *wartegID + "&menu_type_id=" + *menuTypeID + "&menu_name=" + *menuName)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	var responseObject model.ResponseMenuList

	json.Unmarshal(body, &responseObject)

	statusCode := responseObject.StatusCode

	if statusCode != 200 {
		return nil, errors.New(responseObject.Message)
	}

	var menus []*model.MenuList

	for i := 0; i < len(responseObject.Data); i++ {
		menu := &model.MenuList{
			MenuID:   responseObject.Data[i].MenuId,
			Name:     responseObject.Data[i].MenuName,
			Price:    responseObject.Data[i].MenuPrice,
			TypeName: responseObject.Data[i].MenuTypeName,
			WartegID: responseObject.Data[i].WartegId,
		}

		menus = append(menus, menu)
	}
	return menus, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
