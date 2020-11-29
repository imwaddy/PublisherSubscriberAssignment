package services

import (
	"PublisherSubscriberAssignment/ConsumerModules/helpers/dbhelpers"
	"PublisherSubscriberAssignment/ConsumerModules/helpers/loghelpers"
	"PublisherSubscriberAssignment/ConsumerModules/models"
	"errors"
)

// SaveHotelData for saving hotel , room and rateplan object(s)
func SaveHotelData(received *models.IncomingMessage) error {

	// get the connection
	connection, err := dbhelpers.GetConnection()
	if err != nil {
		loghelpers.Error("Error while getting connection: ", err)
		return err
	}

	// if not received anything then throw error
	if received == nil {
		loghelpers.Error("Invalid data received: ", received)
		return errors.New("Invalid data received")
	}

	// looping through incoming objects
	for _, offer := range received.Offers {

		// save hotel object
		err := connection.Save(&offer.Hotel).Error
		if err != nil {
			loghelpers.Error("Error while saving Hotel: ", err)
			return err
		}

		// save room object
		err = connection.Save(&offer.Room).Error
		if err != nil {
			loghelpers.Error("Error while saving Room: ", err)
			return err
		}

		// save rate plan object
		err = connection.Save(&offer.RatePlan).Error
		if err != nil {
			loghelpers.Error("Error while saving RatePlan: ", err)
			return err
		}
	}

	return nil
}
