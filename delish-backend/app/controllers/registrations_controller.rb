class RegistrationsController < ApplicationController
  # Create a new user
  def create
    user = User.new(sign_up_params)

    if user.save
    token = user.generate_jwt
      render json: token.to_json
    else
      render json: { errors: { 'email or password' => ['is invalid'] } }, status: :unprocessable_entity
    end
  end
end
