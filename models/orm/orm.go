package orm

import (
	cgorm "github.com/rasyidmm/RumahMakan/db"
	"gorm.io/gorm"
)

type (
	DBFunc func(tx *gorm.DB) error
)
func Create(v interface{})error  {
	return WithinTransaction(func(tx *gorm.DB)(err error){

		if err = tx.Create(v).Error;err != nil{
			tx.Rollback() // rollback
			return err
		}
		return err
	})
}
func Save(v interface{})error  {
	return WithinTransaction(func(tx *gorm.DB)(err error){

		if err = tx.Save(v).Error;err != nil{
			tx.Rollback() // rollback
			return err
		}
		return err
	})
}

func Delete(v interface{}) error {
	return WithinTransaction(func(tx *gorm.DB) (err error) {
		// check new object
		if err = tx.Delete(v).Error; err != nil {
			tx.Rollback() // rollback
			return err
		}
		return err
	})
}

func FindOneByID(v interface{}, id string) (err error) {
	return WithinTransaction(func(tx *gorm.DB) error {
		if err = tx.First(v, id).Error; err != nil {
			tx.Rollback() // rollback db transaction
			return err
		}
		return err
	})
}

func FindAll(v interface{}) (err error) {
	return WithinTransaction(func(tx *gorm.DB) error {
		if err = tx.Find(v).Error; err != nil {
			tx.Rollback() // rollback db transaction
			return err
		}
		return err
	})
}
//func FindOneByQuery(v interface{}, params map[string]interface{}) (err error) {
//	return WithinTransaction(func(tx *gorm.DB) error {
//		if err = tx.Where(params).Last(v).Error; err != nil {
//			tx.Rollback() // rollback db transaction
//			return err
//		}
//		return err
//	})
//}
//
//func FindByQuery(v interface{}, params map[string]interface{}) (err error) {
//	return WithinTransaction(func(tx *gorm.DB) error {
//		if err = tx.Where(params).Find(v).Error; err != nil {
//			tx.Rollback() // rollback db transaction
//			return err
//		}
//		return err
//	})
//}


func WithinTransaction(fn DBFunc) (err error) {
	tx := cgorm.Init()// start db transaction
	defer tx.Commit()
	err = fn(tx)
	// close db transaction
	return err
}