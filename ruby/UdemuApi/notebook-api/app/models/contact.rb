class Contact < ApplicationRecord
  def author
    "Rod"
  end

  def as_json(options={})
    super(
        root: true,
        methods: :author
    )
  end
end
