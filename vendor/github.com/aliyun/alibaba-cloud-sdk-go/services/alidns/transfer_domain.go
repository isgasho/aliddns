package alidns

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// TransferDomain invokes the alidns.TransferDomain API synchronously
// api document: https://help.aliyun.com/api/alidns/transferdomain.html
func (client *Client) TransferDomain(request *TransferDomainRequest) (response *TransferDomainResponse, err error) {
	response = CreateTransferDomainResponse()
	err = client.DoAction(request, response)
	return
}

// TransferDomainWithChan invokes the alidns.TransferDomain API asynchronously
// api document: https://help.aliyun.com/api/alidns/transferdomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TransferDomainWithChan(request *TransferDomainRequest) (<-chan *TransferDomainResponse, <-chan error) {
	responseChan := make(chan *TransferDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TransferDomain(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// TransferDomainWithCallback invokes the alidns.TransferDomain API asynchronously
// api document: https://help.aliyun.com/api/alidns/transferdomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TransferDomainWithCallback(request *TransferDomainRequest, callback func(response *TransferDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TransferDomainResponse
		var err error
		defer close(result)
		response, err = client.TransferDomain(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// TransferDomainRequest is the request struct for api TransferDomain
type TransferDomainRequest struct {
	*requests.RpcRequest
	DomainNames  string           `position:"Query" name:"DomainNames"`
	Remark       string           `position:"Query" name:"Remark"`
	TargetUserId requests.Integer `position:"Query" name:"TargetUserId"`
	UserClientIp string           `position:"Query" name:"UserClientIp"`
	Lang         string           `position:"Query" name:"Lang"`
}

// TransferDomainResponse is the response struct for api TransferDomain
type TransferDomainResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskId    int64  `json:"TaskId" xml:"TaskId"`
}

// CreateTransferDomainRequest creates a request to invoke TransferDomain API
func CreateTransferDomainRequest() (request *TransferDomainRequest) {
	request = &TransferDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "TransferDomain", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateTransferDomainResponse creates a response to parse from TransferDomain response
func CreateTransferDomainResponse() (response *TransferDomainResponse) {
	response = &TransferDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}