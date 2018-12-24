create_table :user_tags, force: :cascade do |t|
  t.belongs_to :user, null: false
  t.belongs_to :tag, null: false
  t.timestamps

  t.index [:user_id, :tag_id], unique: true
end
