create table suppliers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    code VARCHAR(255) NOT NULL UNIQUE,
    balance INT NOT NULL DEFAULT 0,
    status INT NOT NULL DEFAULT 1,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

create table supplier_recharges (
    id INT AUTO_INCREMENT PRIMARY KEY,
    supplier_id INT NOT NULL,
    supplier_name VARCHAR(255) NOT NULL,
    supplier_code VARCHAR(255) NOT NULL,
    amount INT NOT NULL DEFAULT 0,
    `status` INT NOT NULL DEFAULT 1,
    apply_user_id INT NOT NULL COMMENT '申请用户ID',
    apply_user_name VARCHAR(255) NOT NULL COMMENT '申请用户',
    audit_user_id INT NOT NULL COMMENT '审核用户ID',
    audit_user_name VARCHAR(255) NOT NULL COMMENT '审核用户',
     image_url  varchar(255) NOT NULL COMMENT '审核图片',
    remark VARCHAR(255) COMMENT '审核备注',
    pass_at DATETIME COMMENT '通过时间',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

create table supplier_products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL,
    code VARCHAR(255) NOT NULL UNIQUE,
    supplier_id INT NOT NULL,
    supplier_name VARCHAR(255) NOT NULL,
    supplier_code VARCHAR(255) NOT NULL,
    face_price INT NOT NULL DEFAULT 0,
    price INT NOT NULL DEFAULT 0,
    `status` INT NOT NULL DEFAULT 1,
    spec_id INT NOT NULL,
    sku_id INT NOT NULL,
    brand_id INT NOT NULL;
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

create table skus (
    id INT AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL UNIQUE,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

create table brands (
    id INT AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL UNIQUE,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

create table specs (
    id INT AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL UNIQUE,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


create table channels (
    id INT AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL,
    app_id CHAR(36) NOT NULL UNIQUE,
    secret_key CHAR(32) NOT NULL,
    white_list VARCHAR(255) NOT NULL,
    `status` INT NOT NULL DEFAULT 1,
    balance INT NOT NULL DEFAULT 0,
    credit_limit INT NOT NULL DEFAULT 0 COMMENT 'credit limit 授信',
    credit_balance INT NOT NULL DEFAULT 0 COMMENT 'credit balance 授信余额',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


create table projects (
    id INT AUTO_INCREMENT PRIMARY KEY,
    channel_id INT NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `status` INT NOT NULL DEFAULT 1,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



create table project_products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL,
    `status` INT NOT NULL DEFAULT 1,
    project_id INT NOT NULL,
    brand_id INT NOT NULL,
    spec_id INT NOT NULL,
    sku_id INT NOT NULL,
    face_price INT NOT NULL DEFAULT 0,
    price INT NOT NULL DEFAULT 0,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `version` INT NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE channel_supplier_products (
  id INT PRIMARY KEY AUTO_INCREMENT,
  channel_product_id INT NOT NULL,
  supplier_product_id INT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);