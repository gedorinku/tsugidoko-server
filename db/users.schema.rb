create_table :users, force: :cascade do |t|
  t.string :name, null: false
  t.string :password_digest, null: false
  t.timestamps

  t.index [:name], unique: true
end
