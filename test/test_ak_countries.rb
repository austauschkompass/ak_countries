# frozen_string_literal: true

require "test_helper"

class TestAkCountries < Minitest::Test
  def test_that_it_has_a_version_number
    refute_nil ::AkCountries::VERSION
  end

  def test_it_localizes_country_name
    assert_equal 'UK', ::AkCountries.find_subdivision_name('GB')
    assert_equal 'USA', ::AkCountries.find_subdivision_name('US')
  end

  def test_it_localizes_subdivision_name
    assert_equal 'England (UK)', ::AkCountries.find_subdivision_name('GB-ENG')
  end
end
