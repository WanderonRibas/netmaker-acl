CREATE TABLE user_network_acls (
    user_name TEXT NOT NULL,
    net_id    TEXT NOT NULL,
    PRIMARY KEY (user_name, net_id),
    FOREIGN KEY (user_name) REFERENCES users(user_name) ON DELETE CASCADE,
    FOREIGN KEY (net_id)    REFERENCES networks(netid)   ON DELETE CASCADE
);