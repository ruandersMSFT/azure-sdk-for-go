// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// AutoscaleSettingResourceCollectionPager provides iteration over AutoscaleSettingResourceCollection pages.
type AutoscaleSettingResourceCollectionPager interface {
	// NextPage returns true if the pager advanced to the next page.
	// Returns false if there are no more pages or an error occurred.
	NextPage(context.Context) bool

	// Page returns the current AutoscaleSettingResourceCollectionResponse.
	PageResponse() *AutoscaleSettingResourceCollectionResponse

	// Err returns the last error encountered while paging.
	Err() error
}

type autoscaleSettingResourceCollectionCreateRequest func(context.Context) (*azcore.Request, error)

type autoscaleSettingResourceCollectionHandleError func(*azcore.Response) error

type autoscaleSettingResourceCollectionHandleResponse func(*azcore.Response) (*AutoscaleSettingResourceCollectionResponse, error)

type autoscaleSettingResourceCollectionAdvancePage func(context.Context, *AutoscaleSettingResourceCollectionResponse) (*azcore.Request, error)

type autoscaleSettingResourceCollectionPager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// creates the initial request (non-LRO case)
	requester autoscaleSettingResourceCollectionCreateRequest
	// callback for handling response errors
	errorer autoscaleSettingResourceCollectionHandleError
	// callback for handling the HTTP response
	responder autoscaleSettingResourceCollectionHandleResponse
	// callback for advancing to the next page
	advancer autoscaleSettingResourceCollectionAdvancePage
	// contains the current response
	current *AutoscaleSettingResourceCollectionResponse
	// status codes for successful retrieval
	statusCodes []int
	// any error encountered
	err error
}

func (p *autoscaleSettingResourceCollectionPager) Err() error {
	return p.err
}

func (p *autoscaleSettingResourceCollectionPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if p.current != nil {
		if p.current.AutoscaleSettingResourceCollection.NextLink == nil || len(*p.current.AutoscaleSettingResourceCollection.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.pipeline.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(p.statusCodes...) {
		p.err = p.errorer(resp)
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *autoscaleSettingResourceCollectionPager) PageResponse() *AutoscaleSettingResourceCollectionResponse {
	return p.current
}

// EventDataCollectionPager provides iteration over EventDataCollection pages.
type EventDataCollectionPager interface {
	// NextPage returns true if the pager advanced to the next page.
	// Returns false if there are no more pages or an error occurred.
	NextPage(context.Context) bool

	// Page returns the current EventDataCollectionResponse.
	PageResponse() *EventDataCollectionResponse

	// Err returns the last error encountered while paging.
	Err() error
}

type eventDataCollectionCreateRequest func(context.Context) (*azcore.Request, error)

type eventDataCollectionHandleError func(*azcore.Response) error

type eventDataCollectionHandleResponse func(*azcore.Response) (*EventDataCollectionResponse, error)

type eventDataCollectionAdvancePage func(context.Context, *EventDataCollectionResponse) (*azcore.Request, error)

type eventDataCollectionPager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// creates the initial request (non-LRO case)
	requester eventDataCollectionCreateRequest
	// callback for handling response errors
	errorer eventDataCollectionHandleError
	// callback for handling the HTTP response
	responder eventDataCollectionHandleResponse
	// callback for advancing to the next page
	advancer eventDataCollectionAdvancePage
	// contains the current response
	current *EventDataCollectionResponse
	// status codes for successful retrieval
	statusCodes []int
	// any error encountered
	err error
}

func (p *eventDataCollectionPager) Err() error {
	return p.err
}

func (p *eventDataCollectionPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if p.current != nil {
		if p.current.EventDataCollection.NextLink == nil || len(*p.current.EventDataCollection.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.pipeline.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(p.statusCodes...) {
		p.err = p.errorer(resp)
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *eventDataCollectionPager) PageResponse() *EventDataCollectionResponse {
	return p.current
}