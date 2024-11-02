CREATE TABLE blog_post (
                           pid	int,
                           title	VARCHAR(512),
                           slug	VARCHAR(512),
                           content	VARCHAR(512),
                           markdown 	VARCHAR(512),
                           category_id	int,
                           user_id	int,
                           view_count	int,
                           type	int,
                           create_at	Datetime,
                           update_at	Datetime
);