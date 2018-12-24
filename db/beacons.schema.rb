create_table :beacons, force: :cascade do |t|
  t.string :bssid, null: false
  t.belongs_to :class_room, null: false
  t.timestamps

  t.index [:bssid], unique: true
end

# foreign_key: trueとしても外部キー制約がつかなくて謎
add_foreign_key :beacons, :class_rooms
