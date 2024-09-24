-- +goose Up
-- +goose StatementBegin
CREATE TABLE checklist_items (
  id SERIAL PRIMARY KEY,
  checklist_id INT NOT NULL,
  item_name TEXT NOT NULL,
  is_completed BOOLEAN DEFAULT FALSE,
  CONSTRAINT fk_checklist_id FOREIGN KEY (checklist_id) REFERENCES checklists(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE checklist_items;
-- +goose StatementEnd
