# frozen_string_literal: true

require_relative "ak_countries/version"

require 'json'

module AkCountries
  class Error < StandardError; end

  def self.find_subdivision_name(cc)
    cc_only, subdivision = cc.split('-')

    country = countries.find do |entry|
      entry['alpha2'] == cc_only
    end

    return if country.nil?

    unless subdivision.nil?
      subdivisions = country['subdivisions']

      return if subdivisions.nil?

      entry = subdivisions.find do |div|
        div['code'] == subdivision
      end

      "#{entry['name']} (#{country['name']})" unless entry.nil?
    else
      country['name']
    end
  end

  private

  def self.countries
    @countries ||= JSON.load(File.open(File.join(__dir__, "..", "countries.json")))
  end
end
