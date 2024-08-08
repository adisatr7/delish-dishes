# typed: true
class ApplicationController < ActionController::API
  respond_to :json
  before_action :process_token

  private
  # Check if the user is authenticated
  def authenticate_user!(options = {})
    head :unauthorized unless signed_in?
  end

  # Process the token from the request headers
  def process_token
    if request.headers['Authorization'].present?
      begin
        jwt_payload = JWT.decode(request.headers['Authorization'].split(' ')[1].remove('"'), Rails.application.secrets.secret_key_base).first
        @current_user_id = jwt_payload['id']
      rescue JWT::ExpiredSignature, JWT::VerificationError, JWT::DecodeError
        head :unauthorized
      end
    end
  end

  # Check if the user is signed in
  def signed_in?
    @current_user_id.present?
  end

  # Get the current user from the token
  def current_user
    @current_user ||= super || User.find(@current_user_id)
  end
end
