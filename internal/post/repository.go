package post

import (
	"context"

	"gorm.io/gorm"
)

type PostRepositoryContract interface {
	GetPosts(ctx context.Context) (*[]PostAndUser, error)
	GetPost(ctx context.Context, postID int) (*PostAndUser, error)
	CreatePost(ctx context.Context, post *Post) error
	UpdatePost(ctx context.Context, postID int, post *Post) error
	DeletePost(ctx context.Context, postID int) error
}

type DBGorm struct {
	DB *gorm.DB
}

func ProviderPostRepository(DB *gorm.DB) DBGorm {
	return DBGorm{DB: DB}
}

func (db *DBGorm) GetPosts(ctx context.Context) (*[]PostAndUser, error) {
	var posts []PostAndUser

	err := db.DB.Raw("SELECT p.id, p.title, p.content, p.created, p.modified, p.user, u.name, u.username, u.email, u.phone FROM posts p INNER JOIN users u ON p.user = u.id").Scan(&posts).Error
	if err != nil {
		return nil, err
	}

	return &posts, nil
}

// type PostRepository struct {
// 	DB *gorm.DB
// }

// import (
// 	"context"
// 	"database/sql"
// )

// const (
// 	QUERY_GET_POSTS   = "SELECT p.id, p.title, p.content, p.created, p.modified, p.user, u.name, u.username, u.email, u.phone FROM posts p INNER JOIN users u ON p.user = u.id"
// 	QUERY_GET_POST    = "SELECT p.id, p.title, p.content, p.created, p.modified, p.user, u.name, u.username, u.email, u.phone FROM posts p INNER JOIN users u ON p.user = u.id WHERE p.id = ?"
// 	QUERY_CREATE_POST = "INSERT INTO posts (title, content, created, modified, user) VALUES (?, ?, ?, ?, ?)"
// 	QUERY_UPDATE_POST = "UPDATE posts SET title = ?, content = ?, modified = ? WHERE id = ?"
// 	QUERY_DELETE_POST = "DELETE FROM posts WHERE id = ?"
// )

// type mariaDBRepository struct {
// 	mariaDB *sql.DB
// }

// func NewPostRepository(mariaDBConnection *sql.DB) PostRepository {
// 	return &mariaDBRepository{mariaDB: mariaDBConnection}
// }

// func (r *mariaDBRepository) GetPosts(ctx context.Context) (*[]PostAndUser, error) {
// 	var posts []PostAndUser

// 	res, err := r.mariaDB.QueryContext(ctx, QUERY_GET_POSTS)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer res.Close()

// 	for res.Next() {
// 		post := &PostAndUser{}
// 		err = res.Scan(
// 			&post.ID,
// 			&post.Title,
// 			&post.Content,
// 			&post.Created,
// 			&post.Modified,
// 			&post.User.ID,
// 			&post.User.Name,
// 			&post.User.Username,
// 			&post.User.Email,
// 			&post.User.Phone,
// 			&post.User.Created,
// 			&post.User.Modified,
// 		)
// 		if err != nil && err == sql.ErrNoRows {
// 			return nil, nil
// 		}
// 		if err != nil {
// 			return nil, err
// 		}

// 		posts = append(posts, *post)
// 	}

// 	return &posts, nil

// }

// func (r *mariaDBRepository) GetPost(ctx context.Context, postID int) (*PostAndUser, error) {
// 	post := &PostAndUser{}

// 	stmt, err := r.mariaDB.PrepareContext(ctx, QUERY_GET_POST)
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer stmt.Close()

// 	err = stmt.QueryRowContext(ctx, postID).Scan(
// 		&post.ID,
// 		&post.Title,
// 		&post.Content,
// 		&post.Created,
// 		&post.Modified,
// 		&post.User.ID,
// 		&post.User.Name,
// 		&post.User.Username,
// 		&post.User.Email,
// 		&post.User.Phone,
// 		&post.User.Created,
// 		&post.User.Modified,
// 	)

// 	if err != nil && err == sql.ErrNoRows {
// 		return nil, nil
// 	}
// 	if err != nil {
// 		return nil, err
// 	}

// 	return post, nil
// }

// func (r *mariaDBRepository) CreatePost(ctx context.Context, post *Post) error {
// 	stmt, err := r.mariaDB.PrepareContext(ctx, QUERY_CREATE_POST)
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.ExecContext(ctx, post.Title, post.Content, post.Created, post.Modified, post.User)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (r *mariaDBRepository) UpdatePost(ctx context.Context, postID int, post *Post) error {
// 	stmt, err := r.mariaDB.PrepareContext(ctx, QUERY_UPDATE_POST)
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.ExecContext(ctx, post.Title, post.Content, post.Modified, postID)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (r *mariaDBRepository) DeletePost(ctx context.Context, postID int) error {
// 	stmt, err := r.mariaDB.PrepareContext(ctx, QUERY_DELETE_POST)
// 	if err != nil {
// 		return err
// 	}

// 	defer stmt.Close()

// 	_, err = stmt.ExecContext(ctx, postID)
// 	if err != nil {
// 		return err
// 	}

// 	return nil

// }
