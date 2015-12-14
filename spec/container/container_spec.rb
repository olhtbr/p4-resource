require 'spec_helper'

describe docker_image('olhtbr/p4-resource:latest') do
  it { should exist }
end
