package spine

import (
	"errors"
	"fmt"

	"github.com/enbility/spine-go/api"
	"github.com/enbility/spine-go/model"
)

func (r *NodeManagementImpl) RequestDestinationListData(remoteDeviceAddress *model.AddressDeviceType, sender api.Sender) (*model.MsgCounterType, *model.ErrorType) {
	return nil, model.NewErrorTypeFromString("Not implemented")
}

func (r *NodeManagementImpl) processReadDestinationListData(featureRemote api.FeatureRemote, requestHeader *model.HeaderType) error {
	data := []model.NodeManagementDestinationDataType{
		r.Device().DestinationData(),
	}
	// add other remote devices here

	cmd := model.CmdType{
		NodeManagementDestinationListData: &model.NodeManagementDestinationListDataType{
			NodeManagementDestinationData: data,
		},
	}

	return featureRemote.Sender().Reply(requestHeader, r.Address(), cmd)
}

func (r *NodeManagementImpl) processReplyDestinationListData(message *api.Message, data model.NodeManagementDestinationListDataType) error {
	return errors.New("Not implemented")
}

func (r *NodeManagementImpl) handleMsgDestinationListData(message *api.Message, data *model.NodeManagementDestinationListDataType) error {
	switch message.CmdClassifier {
	case model.CmdClassifierTypeRead:
		return r.processReadDestinationListData(message.FeatureRemote, message.RequestHeader)

	case model.CmdClassifierTypeReply:
		if err := r.pendingRequests.Remove(message.DeviceRemote.Ski(), *message.RequestHeader.MsgCounterReference); err != nil {
			return errors.New(err.String())
		}
		return r.processReplyDestinationListData(message, *data)

	case model.CmdClassifierTypeNotify:
		return r.processReplyDestinationListData(message, *data)

	default:
		return fmt.Errorf("nodemanagement.handleMsgDestinationListData: NodeManagementDestinationListDataType CmdClassifierType not implemented: %s", message.CmdClassifier)
	}
}