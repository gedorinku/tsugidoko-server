create_table :user_positions, force: :cascade do |t|
  t.belongs_to :user, null: false
  t.belongs_to :class_room, null: false
  t.timestamps

  t.index [:user_id, :class_room_id], unique: true
end

# foreign_key: trueとしても外部キー制約がつかなくて謎
add_foreign_key :user_positions, :users
add_foreign_key :user_positions, :class_rooms
