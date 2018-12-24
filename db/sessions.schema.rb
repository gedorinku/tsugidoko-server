create_table :sessions, force: :cascade do |t|
    t.string :secret_key, null: false
    t.belongs_to :user
    t.timestamps
  end
  