# typed: true
require 'securerandom'

class ApplicationRecord < ActiveRecord::Base
  primary_abstract_class
  before_create :assign_uuid

  private
  # Assign a UUID to the record's ID
  def assign_uuid
    self.id = SecureRandom.uuid
  end
end
