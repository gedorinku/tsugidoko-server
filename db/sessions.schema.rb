create_table :sessions, force: :cascade do |t|
  t.string :secret_key, null: false
  t.belongs_to :user, null: false
  t.timestamps

  t.index [:secret_key], unique: true
end

# foreign_key: trueとしても外部キー制約がつかなくて謎
add_foreign_key :sessions, :users
