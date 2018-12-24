create_table :tags, force: :cascade do |t|
  t.string :name, null: false
  t.timestamps

  t.index [:name], unique: true
end
