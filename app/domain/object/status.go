package object

type Status struct {
	ID int64 `json:"id"`

	// AccountID is the account id
	AccountID AccountID `json:"-" db:"account_id"`

	Account Account `json:"account,omitempty"`

	//Content is the status content
	Content string `json:"content"`

	// CreateAt is the time when the status was created
	CreateAt DateTime `json:"create_at,omitempty" db:"create_at"`

	// DeletedAt is the time when the status was deleted
	DeletedAt int64 `json:"-" db:"deleted_at"`
}
