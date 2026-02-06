package service

// type ChannelService struct {
// }

// func (s *ChannelService) CreateChannel(channel *dto.Channel) error {
// 	return repository.CreateChannel(&repository.Channel{
// 		Name:        channel.Name,
// 		WhiteList:   channel.WhiteList,
// 		Status:      channel.Status,
// 		Balance:     0,
// 		CreditLimit: channel.CreditLimit,
// 	})
// }

// func (s *ChannelService) GetChannelByID(id int) (*dto.Channel, error) {
// 	channel, err := repository.GetChannelByID(id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &dto.Channel{
// 		ID:          channel.ID,
// 		Name:        channel.Name,
// 		AppID:       channel.AppID,
// 		SecretKey:   channel.SecretKey,
// 		WhiteList:   channel.WhiteList,
// 		Status:      channel.Status,
// 		Balance:     channel.Balance,
// 		CreditLimit: channel.CreditLimit,
// 		CreatedAt:   channel.CreatedAt,
// 	}, nil
// }
