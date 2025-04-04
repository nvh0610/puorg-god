CREATE TABLE users (
   id          INT PRIMARY KEY AUTO_INCREMENT,
   username    VARCHAR(50) UNIQUE NOT NULL,
   email       VARCHAR(100) UNIQUE NOT NULL,
   password    VARCHAR(255) NOT NULL, -- Mật khẩu đã hash
   role        ENUM('admin', 'user') NOT NULL DEFAULT 'user', -- Phân quyền
   created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE recipes ( -- Công thức
     id          INT PRIMARY KEY AUTO_INCREMENT,
     title       VARCHAR(255) NOT NULL,
     description TEXT NOT NULL,
     image_url   VARCHAR(255), -- Ảnh minh họa
     cuisine     VARCHAR(100), -- Loại ẩm thực (Việt, Hàn, Âu, v.v.)
     created_by  INT NOT NULL,
     created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE ingredients ( -- Nguyên liệu
         id         INT PRIMARY KEY AUTO_INCREMENT,
         name       VARCHAR(100) UNIQUE NOT NULL,
         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
         updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE recipe_ingredients ( -- Công thức + nguyên liệu
                id            INT PRIMARY KEY AUTO_INCREMENT,
                recipe_id     INT NOT NULL,
                ingredient_id INT NOT NULL,
                quantity      VARCHAR(100) NOT NULL, -- Ví dụ: "200g", "2 muỗng"
                created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                updated_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                FOREIGN KEY (recipe_id) REFERENCES recipes(id) ON DELETE CASCADE,
                FOREIGN KEY (ingredient_id) REFERENCES ingredients(id) ON DELETE CASCADE
);

CREATE TABLE instructions ( -- Hướng dẫn từng bước
          id        INT PRIMARY KEY AUTO_INCREMENT,
          recipe_id INT NOT NULL,
          step      INT NOT NULL, -- Bước thứ mấy
          content   TEXT NOT NULL, -- Mô tả bước thực hiện công thức có cả hình ảnh
          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
          FOREIGN KEY (recipe_id) REFERENCES recipes(id) ON DELETE CASCADE
);
