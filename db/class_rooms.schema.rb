create_table :class_rooms, force: :cascade do |t|
  t.string :name, null: false
  t.float :latitude, null: false
  t.float :longitude, null: false
  t.integer :building_id, null: false
  t.integer :floor, null: false
  t.float :local_x, null: false, default: 0.0
  t.float :local_y, null: false, default: 0.0
  t.timestamps
end
