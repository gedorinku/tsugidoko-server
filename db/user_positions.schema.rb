create_table :user_positions, force: :cascade do |t|
  t.belongs_to :user, null: false, index: { unique: true }, foreign_key: true
  t.belongs_to :class_room, null: false
  t.timestamps
end

# foreign_key: trueとしても外部キー制約がつかなくて謎
add_foreign_key :user_positions, :users
add_foreign_key :user_positions, :class_rooms
