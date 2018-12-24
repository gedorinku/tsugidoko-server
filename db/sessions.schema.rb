create_table :sessions, force: :cascade do |t|
  t.string :secret_key, null: false
  t.belongs_to :user, null: false
  t.timestamps

  t.index [:secret_key], unique: true
end
