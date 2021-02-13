package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bytes"
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

func (r *mutationResolver) AddMenu(ctx context.Context, input model.InputMenu) (*model.Resdata, error) {
	requestBody, err := json.Marshal(&model.AddMenu{
		MenuName:    input.Name,
		MenuDetail:  input.Detail,
		MenuPicture: input.Picture,
		MenuPrice:   input.Price,
		MenuTypeId:  input.MenuTypeId,
		WartegId:    input.WartegId,
	})

	client := http.Client{}
	req, _ := http.NewRequest("POST", os.Getenv("MENU_SVC_URL")+"menu", bytes.NewBuffer(requestBody))
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	var responseObject model.ResponseMenu

	json.Unmarshal(body, &responseObject)

	status := responseObject.StatusCode

	var addmenures *model.Resdata

	addmenures = &model.Resdata{
		StatusCode: status,
		Message:    responseObject.Message,
	}

	return addmenures, nil
}

func (r *mutationResolver) UpdateMenu(ctx context.Context, input model.InputMenu, menuID string) (*model.Resdata, error) {
	requestBody, err := json.Marshal(&model.AddMenu{
		MenuName:    input.Name,
		MenuDetail:  input.Detail,
		MenuPicture: input.Picture,
		MenuPrice:   input.Price,
		MenuTypeId:  input.MenuTypeId,
		WartegId:    input.WartegId,
	})

	client := http.Client{}
	req, _ := http.NewRequest("PUT", os.Getenv("MENU_SVC_URL")+"menu/"+menuID, bytes.NewBuffer(requestBody))
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	var responseObject model.ResponseMenu

	json.Unmarshal(body, &responseObject)

	status := responseObject.StatusCode

	var updmenures *model.Resdata

	updmenures = &model.Resdata{
		StatusCode: status,
		Message:    responseObject.Message,
	}

	return updmenures, nil
}

func (r *mutationResolver) DeleteMenu(ctx context.Context, menuID string) (*model.Resdata, error) {

	client := http.Client{}
	req, _ := http.NewRequest("DELETE", os.Getenv("MENU_SVC_URL")+"menu/"+menuID, bytes.NewBuffer(nil))
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	var responseObject model.ResponseMenu

	json.Unmarshal(body, &responseObject)

	status := responseObject.StatusCode

	var delmenures *model.Resdata

	delmenures = &model.Resdata{
		StatusCode: status,
		Message:    responseObject.Message,
	}

	return delmenures, nil

}

func (r *mutationResolver) AddWarteg(ctx context.Context, input model.InputWarteg) (*model.Resdata, error) {
	requestBody, err := json.Marshal(&model.AddWarteg{
		WartegName:        input.Name,
		WartegDesc:        input.Description,
		WartegAddr:        input.Address,
		WartegContactName: input.ContactName,
		WartegPhone:       input.Phone,
	})

	client := http.Client{}
	req, _ := http.NewRequest("POST", os.Getenv("WARTEG_SVC_URL")+"warteg", bytes.NewBuffer(requestBody))
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	var responseObject model.ResponseWartegAdd

	json.Unmarshal(body, &responseObject)

	status := responseObject.StatusCode

	var addwartegres *model.Resdata

	addwartegres = &model.Resdata{
		StatusCode: status,
		Message:    responseObject.Message,
	}

	return addwartegres, nil
}

func (r *mutationResolver) UpdateWarteg(ctx context.Context, input model.InputWarteg, wartegID string) (*model.Resdata, error) {
	requestBody, err := json.Marshal(&model.AddWarteg{
		WartegName:        input.Name,
		WartegDesc:        input.Description,
		WartegAddr:        input.Address,
		WartegContactName: input.ContactName,
		WartegPhone:       input.Phone,
	})

	client := http.Client{}
	req, _ := http.NewRequest("PUT", os.Getenv("WARTEG_SVC_URL")+"warteg/"+wartegID, bytes.NewBuffer(requestBody))
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	var responseObject model.ResponseWartegAdd

	json.Unmarshal(body, &responseObject)

	status := responseObject.StatusCode

	var updwartegres *model.Resdata

	updwartegres = &model.Resdata{
		StatusCode: status,
		Message:    responseObject.Message,
	}

	return updwartegres, nil
}

func (r *mutationResolver) DeleteWarteg(ctx context.Context, wartegID string) (*model.Resdata, error) {
	client := http.Client{}
	req, _ := http.NewRequest("DELETE", os.Getenv("WARTEG_SVC_URL")+"warteg/"+wartegID, bytes.NewBuffer(nil))
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	var responseObject model.ResponseWartegAdd

	json.Unmarshal(body, &responseObject)

	status := responseObject.StatusCode

	var delwartegres *model.Resdata

	delwartegres = &model.Resdata{
		StatusCode: status,
		Message:    responseObject.Message,
	}

	return delwartegres, nil
}

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

func (r *queryResolver) Wartegdetail(ctx context.Context, wartegID string) (*model.WartegDetail, error) {
	resp, err := http.Get(os.Getenv("WARTEG_SVC_URL") + "warteg/" + wartegID)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	var responseObject model.ResponseWartegDetail

	json.Unmarshal(body, &responseObject)

	statusCode := responseObject.StatusCode

	if statusCode != 200 {
		return nil, errors.New(responseObject.Message)
	}

	wartegDetail := &model.WartegDetail{
		WartegId:    responseObject.Data.WartegId,
		Name:        responseObject.Data.WartegName,
		Desc:        responseObject.Data.WartegDesc,
		Address:     responseObject.Data.WartegAddress,
		ContactName: responseObject.Data.WartegContactName,
		Phone:       responseObject.Data.WartegPhone,
	}
	return wartegDetail, nil
}

func (r *queryResolver) Warteglist(ctx context.Context, wartegName *string) ([]*model.WartegList, error) {
	resp, err := http.Get(os.Getenv("WARTEG_SVC_URL") + "wartegs/list?warteg_name=" + *wartegName)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	var responseObject model.ResponseWartegList

	json.Unmarshal(body, &responseObject)

	statusCode := responseObject.StatusCode

	if statusCode != 200 {
		return nil, errors.New(responseObject.Message)
	}

	var wartegs []*model.WartegList

	for i := 0; i < len(responseObject.Data); i++ {
		warteg := &model.WartegList{
			WartegId: responseObject.Data[i].WartegId,
			Name:     responseObject.Data[i].WartegName,
			Address:  responseObject.Data[i].WartegAddr,
		}

		wartegs = append(wartegs, warteg)
	}
	return wartegs, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
