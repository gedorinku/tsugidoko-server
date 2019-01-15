create_table :buildings, force: :cascade do |t|
  t.string :name, null: false
  t.float :latitude, null: false
  t.float :longitude, null: false
  t.timestamps
end
