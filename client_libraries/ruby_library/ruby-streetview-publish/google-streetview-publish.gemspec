# -*- ruby -*-
# encoding: utf-8

Gem::Specification.new do |gem|
  gem.name          = "google-streetview-publish"
  gem.version       = "0.6.8"

  gem.authors       = ["Google Inc"]
  gem.email         = "googleapis-packages@google.com"
  gem.description   = "google-streetview-publish is the official library for Street View Publish API."
  gem.summary       = "API Client library for Street View Publish API"
  gem.homepage      = "https://github.com/googleapis/googleapis"
  gem.license       = "Apache-2.0"

  gem.platform      = Gem::Platform::RUBY

  gem.files         = `git ls-files -- lib/*`.split("\n") +
                      ["README.md", "LICENSE", ".yardopts"]
  gem.require_paths = ["lib"]

  gem.required_ruby_version = ">= 2.0.0"

  gem.add_dependency "google-gax", "~> 0.8.0"
  gem.add_dependency "googleapis-common-protos", "~> 1.3.1"

  gem.add_development_dependency "minitest", "~> 5.10"
  gem.add_development_dependency "rubocop", "<= 0.35.1"
  gem.add_development_dependency "simplecov", "~> 0.9"
end
