create_table :class_room_tags, force: :cascade do |t|
  t.belongs_to :class_room, null: false
  t.belongs_to :tag, null: false
  t.integer :count, null: false, default: 0
  t.timestamps

  t.index [:class_room_id, :tag_id], unique: true
end
  
# foreign_key: trueとしても外部キー制約がつかなくて謎
add_foreign_key :class_room_tags, :class_rooms
add_foreign_key :class_room_tags, :tags
