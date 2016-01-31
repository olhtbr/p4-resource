require 'spec_helper'

describe docker_image('olhtbr/p4-resource:latest') do
  it { should exist }
  its(['Architecture']) { should eq 'amd64' }
  its(['Os']) { should eq 'linux' }
  its(['Config.Entrypoint']) { should include '/init' }
  its(['Config.ExposedPorts']) { should be_nil }
end
