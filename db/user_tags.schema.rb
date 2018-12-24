create_table :user_tags, force: :cascade do |t|
  t.belongs_to :user, null: false
  t.belongs_to :tag, null: false
  t.timestamps

  t.index [:user_id, :tag_id], unique: true
end

# foreign_key: trueとしても外部キー制約がつかなくて謎
add_foreign_key :user_tags, :users
add_foreign_key :user_tags, :tags
